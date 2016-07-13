package commands

import (
	"fmt"
	"os"
    "strings"

	allprojects "github.com/SvenDowideit/gendoc/allprojects"

	"github.com/codegangsta/cli"
)

var logFlag, diffFlag bool

var Status = cli.Command{
	Name:  "status",
	Usage: "status versions from "+allprojects.AllProjectsPath+" file",
	Flags: []cli.Flag{
		cli.BoolFlag{
			Name:        "log",
			Usage:       "Show the commits that are different between checkout and remote",
			Destination: &logFlag,
		},
		cli.BoolFlag{
			Name:        "diff",
			Usage:       "Show the changes that are different between checkout and remote",
			Destination: &diffFlag,
		},
	},
	Action: func(context *cli.Context) error {
		setName, projects, err := allprojects.Load(allprojects.AllProjectsPath)
		if err != nil {
            if os.IsNotExist(err) {
                fmt.Printf("Please run `clone` command first.\n")
            }
			return err
		}
		fmt.Printf("publish-set: %s\n", setName)

        // TODO: confirm what is checked out is the ref from the all-projects
        // TODO: confirm there are no changes - or list them
        for _, p := range *projects {
            // TODO: don't ignore errors.
            fmt.Printf("-- %s (in %s)\n", p.RepoName, *p.Path)
            status(p.RepoName)

            // `gendoc update-yaml` :compare 
            // the sha we have checked out `git rev-parse --verify --quiet HEAD`
            // the head of the branch `git show-ref upstream/master`
            // and the ref sha in the all-projects
            // if they differ, then tell the user they can update the all-projects
            // add a DOIT flag
            // TODO: show isDirty!
            currentSha, err := allprojects.GitResultsIn(p.RepoName, "rev-parse", "--verify", "--quiet", "HEAD")
            currentSha = strings.TrimSpace(currentSha)
            if err != nil {
                fmt.Printf("error: failed to run `git rev-parse --verify --quiet HEAD`\n")
            }
            // TODO: need to use the branch name from ...
            masterBranch := "upstream/master"
            masterSha, err := allprojects.GitResultsIn(p.RepoName, "show-ref", "--hash", masterBranch)
            if err != nil {
                fmt.Printf("error: failed to run `git show-rev %s --hash`\n", masterBranch)
                masterSha = masterBranch
            }
            masterSha = strings.TrimSpace(masterSha)

            differences := false

            allProjectsSHA := p.Ref
            if sha, err := allprojects.GitResultsIn(p.RepoName, "show-ref", "--hash", "upstream/" + p.Ref); err == nil {
                // convert to SHA if ref: master
                allProjectsSHA = strings.TrimSpace(sha)
            }
            if allProjectsSHA != currentSha {
                differences = true
                fmt.Printf("Checkout Sha (%s) NOT the same as ref: in all-projects %s (%s)\n", currentSha, p.Ref, allProjectsSHA)
            }

            if currentSha != masterSha {
                differences = true
                fmt.Printf("Checkout Sha (%s) NOT the same as tip of %s (%s)\n", currentSha, masterBranch, masterSha)
                fmt.Printf("\tConsider a `git reset --hard %s`\n", masterBranch)
            }
            if (logFlag || diffFlag) && differences {
                allprojects.GitIn(p.RepoName, "log", "--first-parent", "--oneline", currentSha+".."+masterSha, *p.Path)
		// TODO 
		//foreach commit in the log, show the diff
		// git diff  a63a059^1..a63a059 docs/
            }
        }
        return nil
	},
}

func status(repoPath string) error {
    // TODO: mmm, how to combine error
    // maybe it should return the list of refs with its HEAD date
    // hopefully the real branch will be most used (eventually we use the Branch: info)
    allprojects.GitIn(repoPath, "branch", "--contains")
    return allprojects.GitIn(repoPath, "tag", "--contains")
}
