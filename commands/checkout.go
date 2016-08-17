package commands

import (
	"fmt"
	"os"
	"strings"

	"github.com/SvenDowideit/gendoc/allprojects"

	"github.com/Sirupsen/logrus"
	"github.com/codegangsta/cli"
)

var fetchFlag, resetFlag bool

var Checkout = cli.Command{
	Name:  "checkout",
	Usage: "checkout versions from " + allprojects.AllProjectsPath + " file",
	Flags: []cli.Flag{
		cli.BoolFlag{
			Name:        "fetch",
			Usage:       "git fetch upstream",
			Destination: &fetchFlag,
		},
		cli.BoolFlag{
			Name:        "reset",
			Usage:       "get reset --hard upstream/<ref>",
			Destination: &resetFlag,
		},
	},
	Action: func(context *cli.Context) error {
		// TODO: checkout what's in the current file - we might be testing a branch
		if context.NArg() == 1 {
			publishSetBranch := context.Args()[0]
			fmt.Printf("Checking out %s %s.\n", allprojects.AllProjectsRepo, publishSetBranch)
			err := checkout(allprojects.AllProjectsRepo, publishSetBranch)
			if err != nil {
				return err
			}
		} else {
			fmt.Printf("Using the docs.docker.com/all-projects.yml as is.\n")

		}

		//TODO need to fetch&reset docs.docker.com, docs-html and docs-src

		setName, projects, err := allprojects.Load(allprojects.AllProjectsPath)
		if err != nil {
			if os.IsNotExist(err) {
				fmt.Printf("Please run `clone` command first.\n")
			}
			return err
		}
		if fetchFlag {
			err := allprojects.GitIn(allprojects.AllProjectsRepo, "fetch", "--all")
			if err != nil {
				return err
			}
			err = allprojects.GitIn(allprojects.AllProjectsRepo, "fetch", "--tag", "upstream")
			if err != nil {
				return err
			}
		}
		fmt.Printf("publish-set: %s\n", setName)

		for _, p := range *projects {
			// TODO: don't ignore errors.
			fmt.Printf("-- %s\n", p.RepoName)
			checkout(p.RepoName, p.Ref)
		}
		return nil
	},
}

//TODO: bail out if there are local commits, or isdirty
func checkout(repoPath, ref string) error {
	if fetchFlag {
		err := allprojects.GitIn(repoPath, "fetch", "--all")
		if err != nil {
			return err
		}
	}

	// exit happy if the sha of HEAD == the SHA that the ref points to (not the sha of the tag)
	msg, err := hasCheckedOutRef(repoPath, ref)
	if err == nil {
		fmt.Printf("Same as all-projects.yml: %s\n", msg)
		return nil
	}

	err = allprojects.GitIn(repoPath, "checkout", ref)
	if err != nil {
		// do a fetch, in case it exists in remote
		err = allprojects.GitIn(repoPath, "fetch", "upstream", ref+":remotes/upstream/"+ref)
		if err != nil {
			// Last resourt, fetch all upstream, and undo depth
			err = allprojects.GitIn(repoPath, "fetch", "--all")
			if err != nil {
				return err
			}
			err = allprojects.GitIn(repoPath, "fetch", "--tag", "upstream")
			if err != nil {
				return err
			}
		}
		err = allprojects.GitIn(repoPath, "checkout", ref)
		if err != nil {
			err = allprojects.GitIn(repoPath, "checkout", "-b", ref, "remotes/upstream/"+ref)
			if err != nil {
				return err
			}
		}
	}
	if resetFlag {
		if _, err := allprojects.GitResultsIn(repoPath, "show-ref", "--hash", "upstream/"+ref); err == nil {
			// its not a SHA, so we can reset
			err = allprojects.GitIn(repoPath, "reset", "--hard", "upstream/"+ref)
			if err != nil {
				return err
			}
		}
	}
	return err
}

func getSHA(repoPath, ref string) (sha string, err error) {
	sha, err = allprojects.GitResultsIn(repoPath, "log", "-1", "--format=%H", ref)
	return strings.TrimSpace(sha), err
}

func hasCheckedOutRef(repoPath, ref string) (string, error) {
	// exit happy if the sha of HEAD == the SHA that the ref points to (not the sha of the tag)
	headSHA, err := getSHA(repoPath, "HEAD")
	if err != nil {
		// if we can't get the SHA of HEAD, we're dead
		return "", err
	}
	logrus.Debugf("compare (%s) to (%s)\n", headSHA, ref)
	if headSHA == ref {
		// the all-projects ref is a SHA
		return fmt.Sprintf("your checkout %s is at %s", headSHA, ref), nil
	}
	// is ref an upstream branch?
	if refSHA, err := getSHA(repoPath, "upstream/"+ref); err == nil {
		logrus.Debugf("compare (%s) to (%s)\n", headSHA, refSHA)
		// if we got that ok, we don't need a checkout / fetch
		if headSHA == refSHA {
			return fmt.Sprintf("your checkout %s is at %s", headSHA, "upstream/"+ref), nil
		}
	}
	// is ref a tag?
	if tagSHA, err := allprojects.GitResultsIn(repoPath, "show-ref", "--hash", "refs/tags/"+ref); err == nil {
		tagSHA = strings.TrimSpace(tagSHA)
		if refSHA, err := getSHA(repoPath, tagSHA); err == nil {
			logrus.Debugf("compare (%s) to (%s)\n", headSHA, refSHA)
			// if we got that ok, we don't need a checkout / fetch
			if headSHA == refSHA {
				return fmt.Sprintf("your checkout %s is at %s", headSHA, "refs/tags/"+ref), nil
			}
		}
	}
	return "", fmt.Errorf("Error: repository not at %s", ref)
}
