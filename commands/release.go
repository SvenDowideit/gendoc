package commands

import (
	"bytes"
	"fmt"
	"strconv"
	"strings"
	"text/template"
	"time"

	"github.com/SvenDowideit/gendoc/allprojects"

	"github.com/Sirupsen/logrus"
	"github.com/blang/semver"
	"github.com/codegangsta/cli"
)

var remoteName = "upstream"
var compareToBranch = remoteName + "/master"
var pushFlag, showFilesFlag, showFutureMilestoneFlag, cherryPickFlag, showLabeledFlag, noisyFlag bool

var Release = cli.Command{
	Name:  "release",
	Usage: "Prepare and ship a docs release.",
	Flags: []cli.Flag{},
	Subcommands: []cli.Command{
		{
			Name:  "prepare",
			Usage: "Prepare docs release tags and branches.",
			Flags: []cli.Flag{
				cli.BoolFlag{
					Name:        "doit",
					Usage:       "list PR's without having a GITHUB_TOKEN",
					Destination: &doitFlag,
				},
				cli.BoolFlag{
					Name:        "noisy",
					Usage:       "tell me about all skipped PR's",
					Destination: &noisyFlag,
				},
				cli.BoolFlag{
					Name:        "cherry-pick",
					Usage:       "Cherry-pick all PR's listed into whatever branch you currently have",
					Destination: &cherryPickFlag,
				},
				cli.BoolFlag{
					Name:        "show-future",
					Usage:       "show PR's that are for a future milestone (compared to the all-project product version)",
					Destination: &showFutureMilestoneFlag,
				},
				cli.BoolFlag{
					Name:        "show-labeled",
					Usage:       "show PR's labeled with 'cherry-picked'",
					Destination: &showLabeledFlag,
				},
				cli.BoolFlag{
					Name:        "files",
					Usage:       "Show the files that changes in each commit",
					Destination: &showFilesFlag,
				},
				cli.StringFlag{
					Name:        "branch",
					Usage:       "Compare all-projects.yml ref's to this branch",
					Value:       "upstream/master",
					Destination: &compareToBranch,
				},
			},
			Action: func(context *cli.Context) error {
				if allprojects.GithubToken == "" {
					if !doitFlag {
						return fmt.Errorf("You have not set the GITHUB_TOKEN env var (or used --ghtoken to set it)\nAdd `--doit` to run anyway - the output may be missing GitHub API specific information.\n")
					}
					fmt.Printf("WARNING: GitHub token not set, your output may be missing PR milestones and labels\n")
				}
				setName, projects, err := allprojects.Load(allprojects.AllProjectsPath)
				if err != nil {
					return err
				}
				fmt.Printf("publish-set: %s\n", setName)
				fmt.Printf("comparing current checkout to %s\n", compareToBranch)

				if context.NArg() > 0 {
					name := context.Args()[0]
					project, err := projects.GetProjectByName(name)
					if err != nil {
						return err
					}
					findDocsPRsNeedingMerge(project)
					return nil
				}

				for _, p := range *projects {
					findDocsPRsNeedingMerge(p)
				}

				return nil
			},
		},
		{
			Name:  "tag",
			Usage: "Check, or create product release tags matching this all-projects.yml.",
			Flags: []cli.Flag{
				cli.BoolFlag{
					Name:        "doit",
					Usage:       "Actualy create and push the tag - without it we only get to see what could have been",
					Destination: &doitFlag,
				},
				cli.BoolFlag{
					Name:        "push",
					Usage:       "Push tags that we didn't create this run",
					Destination: &pushFlag,
				},
				cli.StringFlag{
					Name:        "remote",
					Usage:       "test or push tags to specified remote",
					Value:       "upstream",
					Destination: &remoteName,
				},
			},
			Action: func(context *cli.Context) error {
				setName, projects, err := allprojects.Load(allprojects.AllProjectsPath)
				if err != nil {
					return err
				}
				fmt.Printf("publish-set: %s\n", setName)

				if context.NArg() > 0 {
					name := context.Args()[0]
					project, err := projects.GetProjectByName(name)
					if err != nil {
						return err
					}
					tagProduct(project)
				} else {
					for _, p := range *projects {
						tagProduct(p)
					}
				}

				return nil
			},
		},
		{
			Name:  "release",
			Usage: "Publish docs.",
			Flags: []cli.Flag{},
			Action: func(context *cli.Context) error {
				setName, _, err := allprojects.Load(allprojects.AllProjectsPath)
				if err != nil {
					return err
				}
				fmt.Printf("publish-set: %s\n", setName)

				return fmt.Errorf("i've got nothin")
			},
		},
	},
}

func getCommitDate(repo, ref string) (strDate string, date time.Time, err error) {
	out, err := allprojects.GitResultsIn(repo, "log", "-1", "--format=%cD", ref, "--")
	if err != nil {
		return "", date, fmt.Errorf("ERROR: failed to get date of %s (%s)\nYou may need to run gendoc checkout --fetch", ref, err)
	}
	o := strings.SplitN(out, "\n", 2)
	strDate = strings.TrimSpace(o[0])
	logrus.Debugf("got %s\n", out)
	date, err = time.Parse("Mon, 2 Jan 2006 15:04:05 -0700", strDate)
	if err != nil {
		return "", date, fmt.Errorf("Failed to Parse %s (%s)\n", out, err)
	}
	return strDate, date, nil
}

//TODO: code from checkout
func getSHA(repoPath, ref string) string {
	// TODO: is it an SHA, return it (fast path)
	// is it an upstream branch?
	if refSHA, err := allprojects.GitResultsIn(repoPath, "log", "-1", "--format=%H", "upstream/"+ref); err == nil {
		refSHA = strings.TrimSpace(refSHA)
		return refSHA
	}
	// is it a tag?
	if tagSHA, err := allprojects.GitResultsIn(repoPath, "show-ref", "--hash", "refs/tags/"+ref); err == nil {
		tagSHA = strings.TrimSpace(tagSHA)
		if refSHA, err := allprojects.GitResultsIn(repoPath, "log", "-1", "--format=%H", tagSHA); err == nil {
			refSHA = strings.TrimSpace(refSHA)
			return refSHA
		}
	}
	return ref
}

// tagProduct will check for the exitence of a product tag in that project's repo
// and will check that it matches the commit listed in the all-projects file
// OR will make that tag
// the tag's date will be either today, or ... (can I get it from the docs-html repo?)
func tagProduct(p allprojects.Project) {
	fmt.Printf("## Tag for  %s in %s\n", p.Name, p.RepoName)
	tmpl := template.New("tag")
	tmpl, _ = tmpl.Parse("docs{{with .Version}}-{{.}}{{end}}-{{.Date}}")

	// Get committer date for commit
	committerDate, date, err := getCommitDate(p.RepoName, p.Ref)
	if err != nil {
		fmt.Printf("ERROR: (%s)\n", err)
		return
	}

	type info struct {
		Date    string
		Version string
	}
	i := info{
		Date:    date.Format("2006-01-02"),
		Version: p.Version,
	}

	pRefSHA := getSHA(p.RepoName, p.Ref)

	var doc bytes.Buffer
	_ = tmpl.Execute(&doc, i)
	tag := doc.String()

	// test to see if the tag is already there, and  see that the tag matches what we would have made..
	// and tell the user otherwise.
	// TODO: this code is very similar to what is in `release prepare`
	out, err := allprojects.GitResultsIn(p.RepoName, "log", "--pretty=format:%H\t%s", "-1", tag)
	if err == nil && out != "" {
		localTag := strings.Split(out, "\t")
		localTag[0] = strings.TrimSpace(localTag[0])
		localTag[1] = strings.TrimSpace(localTag[1])
		if localTag[0] == pRefSHA {
			logrus.Debugf("tag already exists locally (%s)\n", tag)
		} else {
			fmt.Printf("tag already exists locally (%s = %s) but differs from all-projects (%s == %s)\n", tag, localTag[0], p.Ref, pRefSHA)
		}
		// "tagname^{}" is the commit SHA the tag is pointing to
		out, err = allprojects.GitResultsIn(p.RepoName, "ls-remote", remoteName, tag+"^{}")
		if err == nil && out != "" {
			remoteTag := strings.Split(out, "\t")
			remoteTag[0] = strings.TrimSpace(remoteTag[0])
			remoteTag[1] = strings.TrimSpace(remoteTag[1])
			if remoteTag[0] == pRefSHA {
				fmt.Printf("- tag %s matches\n", p.Ref)
			} else {
				fmt.Printf("ERROR: tag already exists on remote (%s = %s) but differs from all-projects (%s = %s)\n", remoteTag[1], remoteTag[0], p.Ref, pRefSHA)
				fmt.Printf("TODO: check feet for holes\n")
			}
		} else {
			if pushFlag {
				fmt.Printf("tag exists locally, pushing to remote remote\n")
				err = allprojects.GitIn(p.RepoName, "push", remoteName, tag)
				if err != nil {
					fmt.Printf("Failed to push tag %s (%s)\n", tag, err)
				} else {
					fmt.Printf("OK: found %s on remote %s (%s)\n", tag, remoteName, p.Ref)
				}
			} else {
				fmt.Printf("ERROR: %s exists locally, but not in remote\n", tag)
				fmt.Printf("TODO: add `--push` to the commandline to push to remote\n")
			}
		}
		return
	}

	if !doitFlag {
		fmt.Printf("- Proposed Tag == %s to %s   (add --doit to the command to create and push)\n", tag, p.Ref)
		// TODO test to see if that's the HEAD of the checkout, and do more warning
		return
	}
	// make an annotated tag
	// TODO: set the tag's date
	out, err = allprojects.GitEnvResultsIn(
		[]string{"GIT_COMMITTER_DATE=" + committerDate},
		p.RepoName, "tag", "-a", "-m", "generated tag from history", tag,
		p.Ref,
	)
	if err != nil {
		fmt.Printf("Failed to make local tag %s (%s)\n", tag, err)
		return
	}
	err = allprojects.GitIn(p.RepoName, "push", remoteName, tag)
	if err != nil {
		fmt.Printf("Failed to push tag to %s (%s)\n", remoteName, err)
		return
	}
	fmt.Printf("OK: created %s on remote %s (%s)\n", tag, remoteName, p.Ref)
}

func parseVersion(version string) (semver.Version, error) {
	// We don't do semver properly, so remove any -rc, -alpha etc
	// Some projects use non-version prefixes - so we need to remove those
	version = strings.TrimPrefix(version, "Notary ")
	version = strings.TrimPrefix(version, "Registry/")
	version = strings.SplitN(version, "-", 2)[0]
	pVersion, err := semver.ParseTolerant(version)
	if err != nil {
		logrus.Debugf("versionParse(%s)\n", version)
	}
	return pVersion, err
}

// I think I can't just use
// git log --merges --oneline 93cc2675c8f97e1a30b3bf2dbc287f0295ffc4fa..upstream/master --parents
// becuase that presumes we have a linear history
func findDocsPRsNeedingMerge(p allprojects.Project) {
	fmt.Printf("\n## %s, %s in %s at %s\n", p.Name, p.Version, p.RepoName, p.Ref)
	pVersion, err := parseVersion(p.Version)
	if err != nil {
		if p.Version == "" {
			fmt.Printf("Warning: no version field in all-projects.yml for %s\n", p.Name)
		} else {
			logrus.Debugf("ERROR parsing %s, (%s) %s\n", pVersion, p.Name, err)
		}
	}
	out, _, err, cmd := allprojects.GitScannerIn(p.RepoName, "cherry", "-v", "HEAD", compareToBranch)
	if err != nil {
		fmt.Printf("ERROR %s\n", err)
		return
	}
	err = cmd.Start()
	if err != nil {
		fmt.Printf("ERROR %s\n", err)
		return
	}
	defer cmd.Wait()
	for out.Scan() {
		line := out.Text()
		//logrus.Debugf("%s\n", out.Text())
		// + ffdef1abbd01c2479d02270d919aed9fa40a52e4 use tabwriter in favour of tablewriter
		oneline := strings.SplitN(line, " ", 3)
		if oneline[0] != "+" {
			continue
		}
		//logrus.Debugf("joined: %s\n", strings.Join(oneline, ","))

		// Find the merge commit for it
		// merge commit with PR# is first line of
		// git log --ancestry-path --merges --oneline --reverse c9bf41955c53cf1780e043db2d8887c2cac62429..upstream/master
		// OR per http://stackoverflow.com/questions/8475448/find-merge-commit-which-include-a-specific-commit
		// git rev-list $1..master --ancestry-path | grep -f <(git rev-list $1..master --first-parent) | tail -1
		ancestor, _, err, ancestorCmd := allprojects.GitScannerIn(p.RepoName, "log", "--ancestry-path", "--merges", "--oneline", "--reverse", oneline[1]+".."+compareToBranch)
		if err != nil {
			fmt.Printf("ERROR find merge commit  %s\n", err)
			//continue
		}
		err = ancestorCmd.Start()
		if err != nil {
			fmt.Printf("ERROR find merge commit  %s\n", err)
			//continue
		}
		// 1e176b5 Merge pull request #3592 from stakodiak/fix-privilege-typo
		if !ancestor.Scan() {
			errStr := ""
			if ancestor.Err() != nil {
				errStr = ancestor.Err().Error()
			}
			fmt.Printf("NO merge PR found for (%s) %s\n", line, errStr)
			ancestorCmd.Wait()
			continue
		}
		text := ancestor.Text()
		if text == "" {
			if ancestor.Scan() {
				fmt.Printf("ERROR scan2 (%s) %s\n", line, ancestor.Err())
				ancestorCmd.Wait()
				continue
			}
			text := ancestor.Text()
			fmt.Printf("-- scan ERROR %s\n", text)
			ancestorCmd.Wait()
			continue
		}
		ancestorCmd.Wait()

		logrus.Debugf("test: %s\n", text)
		a := strings.Split(text, " ")
		mergeSHA := a[0]
		mergePR := 0
		mergeBranch := "NOT GitHub"
		if a[1] == "Merge" && a[2] == "pull" && a[3] == "request" && a[5] == "from" {
			// Then this is likely a GitHub PR merge commit
			mergePR, _ = strconv.Atoi(strings.TrimLeft(a[4], "#"))
			mergeBranch = a[6]
		}

		// Find out if there were doc changes..
		// git diff-tree --no-commit-id --name-only -r <sha> <docs-dir>
		files, _, err, filesCmd := allprojects.GitScannerIn(p.RepoName, "diff-tree", "--no-commit-id", "--name-only", "-r", oneline[1], *p.Path)
		if err != nil {
			fmt.Printf("ERROR diff-tree %s\n", err)
			//continue
		}
		err = filesCmd.Start()
		if err != nil {
			fmt.Printf("ERROR diff-tree %s\n", err)
			//continue
		}
		if !files.Scan() {
			// TODO: maybe only do the find merge PR if debug?
			logrus.Debugf("%d (merge commit: %s ) from %s\n", mergePR, mergeSHA, mergeBranch)
			// labels, milestone, _ := allprojects.GetPRInfo(p.Org, p.RepoName, mergePR)
			// logrus.Debugf("\t %s: %s\n", milestone, labels)
			logrus.Debugf("\tNO %s changes in %s %s\n", *p.Path, oneline[1], oneline[2])
			filesCmd.Wait()
			continue
		}

		labels, milestone, err := allprojects.GetPRInfo(p.Org, p.RepoName, mergePR)
		if pVersion.Major > 0 || pVersion.Minor > 0 || pVersion.Patch > 0 {
			mVersion, err := parseVersion(milestone)
			if err != nil {
				if milestone != "" {
					fmt.Printf("ERROR parsing Version(%s) in milestone of PR(%d) %s\n", milestone, mergePR, err)
				}
			} else {
				if !showFutureMilestoneFlag && mVersion.GT(pVersion) {
					if noisyFlag {
						fmt.Printf("Skipping %d due to %s\n", mergePR, milestone)
					}
					continue
				}
			}
		}
		if !showLabeledFlag {
			// if the labels contain process/cherry-picked or process/docs-cherry-picked, skip
			if strings.Contains(labels, "cherry-picked") {
				if noisyFlag {
					fmt.Printf("Skipping %d due to cherry-picked state: %s\n", mergePR, labels)
				}
				continue
			}
			if strings.Contains(labels, "process/cherry-pick") {
				if noisyFlag {
					fmt.Printf("Skipping %d due to code cherry-pick label: %s\n", mergePR, labels)
				}
				continue
			}
		}

		// Last attempt to match - see if there is a cherry-pick -x -m1 commit in the destination repo that was ammended manually
		// get a list of commits from HEAD, see if there's a "Merge pull request #25405 from thaJeztah/fix-api-markdown" like commit
		// if so - its already been picked, we can skip.
		cherryPickSHA, err := findCherryPickCommit(p.RepoName, fmt.Sprintf("%d", mergePR))
		if err == nil {
			if noisyFlag {
				fmt.Printf("Skipping %d due to cherry-picked commit: %s (presumably an ammended cherry-pick compared to %s)\n", mergePR, cherryPickSHA, mergeSHA)
			}
			continue
		}


		////////////////
		// OK - so we've decided to show these PR's
		_, mergeDate, _ := getCommitDate(p.RepoName, mergeSHA)
		fmt.Printf("- PR %d (%s) %s from %s\n", mergePR, mergeSHA, mergeDate.UTC().Format(time.Stamp), mergeBranch)
		if milestone == "" {
			fmt.Printf("- Warning: no version milestone set for PR(%d)\n", mergePR)
		}
		if milestone != "" || labels != "" {
			fmt.Printf("- %s %s\n", milestone, labels)
		}


		fmt.Printf("  - %s changes in %s %s\n", *p.Path, oneline[1], oneline[2])
		if showFilesFlag {
			fmt.Printf("    - %s\n", files.Text())
			for files.Scan() {
				fmt.Printf("  - %s\n", files.Text())
			}
		}
		filesCmd.Wait()
		if cherryPickFlag {
			err = allprojects.GitIn(p.RepoName, "cherry-pick", "-s", "-x", "-m1", mergeSHA)
			if err != nil {
				fmt.Printf("help needed to cherry-pick PR %d (%s) %s\n", mergePR, mergeSHA, err)
				return // the user can restart the process after fixing things up
			} else {
				fmt.Printf("Cherry-picked PR %d (%s)\n", mergePR, mergeSHA)
			}
		}
	}
	if err := out.Err(); err != nil {
		fmt.Printf("ERROR: %s\n", err)
	}
}

// findCherryPickCommit given a PR number, iterate through the commits to see if there is a matching cherry-pick commit
func findCherryPickCommit(repo string, pr string) (commitSHA string, err error) {
	logrus.Debugf("findCherryPickCommit(%s, %s)", repo, pr)
	log, _, err, cmd := allprojects.GitScannerIn(repo, "log", "--oneline", "HEAD")
	if err != nil {
		fmt.Printf("ERROR %s\n", err)
		return "", err
	}
	err = cmd.Start()
	if err != nil {
		fmt.Printf("ERROR %s\n", err)
		return "", err
	}
	defer cmd.Wait()
	for log.Scan() {
		if log.Err() != nil {
			return "", log.Err()
		}
		text := log.Text()
		// TODO: cache this as a hash of PR -> sha
		details := strings.SplitN(text, " ", 2)
		expectedCommitPrefix := fmt.Sprintf("Merge pull request #%s ", pr)
		if strings.HasPrefix(details[1], expectedCommitPrefix) {
			// TODO: can be more careful, and check for "(cherry picked from commit 66671d4ec29d7ccbd991399b8b98705e57b6a3eb)"
			// but really, if there is a merge commit for a PR in the local branch to match any PR in master, we should be good to go
			return details[0], nil
		}
	}
	return "", fmt.Errorf("no merge commit found for PR %s", pr)
}
