package commands

import (
	"fmt"
	"strings"
	"strconv"

	allprojects "github.com/SvenDowideit/gendoc/allprojects"

	"github.com/codegangsta/cli"
	"github.com/Sirupsen/logrus"
)

var Release = cli.Command{
	Name:  "release",
	Usage: "Prepare and ship a docs release.",
	Flags: []cli.Flag{
	},
	Subcommands: []cli.Command{
        {
            Name:  "prepare",
            Usage: "Prepare docs release tags and branches.",
            Flags: []cli.Flag{
            },
            Action: func(context *cli.Context) error {
                setName, projects, err := allprojects.Load(allprojects.AllProjectsPath)
                if err != nil {
                    return err
                }
                fmt.Printf("publish-set: %s\n", setName)
		if context.NArg() > 0 {
			name := context.Args()[0]
			project := projects.GetProjectByName(name)
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
            Name:  "push",
            Usage: "Push docs release tags&branches to all the repos.",
            Flags: []cli.Flag{
            },
            Action: func(context *cli.Context) error {
                setName, _, err := allprojects.Load(allprojects.AllProjectsPath)
                if err != nil {
                    return err
                }
                fmt.Printf("publish-set: %s\n", setName)

                return fmt.Errorf("i've got nothin")
            },
        },
        {
            Name:  "release",
            Usage: "Publish docs.",
            Flags: []cli.Flag{
            },
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


// I think I can't just use 
// git log --merges --oneline 93cc2675c8f97e1a30b3bf2dbc287f0295ffc4fa..upstream/master --parents
// becuase that presumes we have a linear history

func findDocsPRsNeedingMerge(p allprojects.Project) {
			fmt.Printf("-- %s in %s\n", p.Name, p.RepoName)
                	out, _, err := allprojects.GitScannerIn(p.RepoName, "cherry", "-v", p.Ref, "upstream/master")
			if err != nil {
				fmt.Printf("ERROR %s\n", err)
				return
			}
			for out.Scan() {
				line := out.Text()
		 		// fmt.Printf("%s\n", out.Text())
				// + ffdef1abbd01c2479d02270d919aed9fa40a52e4 use tabwriter in favour of tablewriter
				oneline := strings.SplitN(line, " ", 3)
				if oneline[0] != "+" {
					continue
				}
				// Find out if there were doc changes..
				// git diff-tree --no-commit-id --name-only -r <sha> <docs-dir>
                		files, err := allprojects.GitResultsIn(p.RepoName, "diff-tree", "--no-commit-id", "--name-only", "-r", oneline[1], *p.Path)
				if err != nil {
					fmt.Printf("ERROR diff-tree %s\n", err)
					//continue
				}
				// Find the merge commit for it
				// merge commit with PR# is first line of 
				// git log --ancestry-path --merges --oneline --reverse c9bf41955c53cf1780e043db2d8887c2cac62429..upstream/master
				// OR per http://stackoverflow.com/questions/8475448/find-merge-commit-which-include-a-specific-commit
				// git rev-list $1..master --ancestry-path | grep -f <(git rev-list $1..master --first-parent) | tail -1
                		ancestor, _, err := allprojects.GitScannerIn(p.RepoName, "log", "--ancestry-path", "--merges", "--oneline", "--reverse", oneline[1]+"..upstream/master")
				if err != nil {
					fmt.Printf("ERROR find merge commit  %s\n", err)
					//continue
				}
				// 1e176b5 Merge pull request #3592 from stakodiak/fix-privilege-typo
				if !ancestor.Scan() {
					fmt.Printf("ERROR scan (%s) %s\n", line, ancestor.Err())
					continue
				}
				text := ancestor.Text()
				if text == "" {
					if ancestor.Scan() {
						fmt.Printf("ERROR scan2 (%s) %s\n", line, ancestor.Err())
						continue
					}
					text := ancestor.Text()
					fmt.Printf("-- scan err %s\n", text)
					continue
				}
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
				if files == "" {
					// TODO: maybe only do the find merge PR if debug?
					labels, milestone, _ := allprojects.GetPRInfo(p.Org, p.RepoName, mergePR)
					logrus.Debugf("%d (%s) from %s\n", mergePR, mergeSHA, mergeBranch)
					logrus.Debugf("\t %s: %s\n", milestone, labels)
					logrus.Debugf("\tNO %s changes in %s %s\n", *p.Path, oneline[1], oneline[2])
					continue
				}
				
				labels, milestone, err := allprojects.GetPRInfo(p.Org, p.RepoName, mergePR)
				fmt.Printf("%d (%s) from %s\n", mergePR, mergeSHA, mergeBranch)
				fmt.Printf("\t %s %s\n", milestone, labels)
				fmt.Printf("\t%s changes in %s %s\n", *p.Path, oneline[1], oneline[2])
				fmt.Printf("%s\n", files)
			}
			if err := out.Err(); err != nil {
			    fmt.Printf("ERROR: %s\n", err)
			}
}
