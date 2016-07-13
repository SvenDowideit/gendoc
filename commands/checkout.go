package commands

import (
	"fmt"
	"os"

	allprojects "github.com/SvenDowideit/gendoc/allprojects"

	"github.com/codegangsta/cli"
)

var fetchFlag, resetFlag bool

var Checkout = cli.Command{
	Name:  "checkout",
	Usage: "checkout versions from "+allprojects.AllProjectsPath+" file",
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

        //TODO need to fetch&reset docs.docekr.com, docs-html and docs-src

		setName, projects, err := allprojects.Load(allprojects.AllProjectsPath)
		if err != nil {
            if os.IsNotExist(err) {
                fmt.Printf("Please run `clone` command first.\n")
            }
			return err
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
        if _, err := allprojects.GitResultsIn(repoPath, "show-ref", "--hash", "upstream/" + ref); err == nil {
            // its not a SHA, so we should fetch
            err = allprojects.GitIn(repoPath, "fetch", "upstream", ref+":remotes/upstream/"+ref)
            if err != nil {
                return err
            }
        }
    }

    //TODO what if its a tag
    err := allprojects.GitIn(repoPath, "checkout", ref)
    if err != nil {
        // do a fetch, in case it exists in remote
        err = allprojects.GitIn(repoPath, "fetch", "upstream", ref+":remotes/upstream/"+ref)
        if err != nil {
            // Last resourt, fetch all upstream, and undo depth
            err = allprojects.GitIn(repoPath, "fetch", "--all")
            if err != nil {
                return err
            }
            err = allprojects.GitIn(repoPath, "fetch", "--tag")
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
        if _, err := allprojects.GitResultsIn(repoPath, "show-ref", "--hash", "upstream/" + ref); err == nil {
            // its not a SHA, so we can reset
            err = allprojects.GitIn(repoPath, "reset", "--hard", "upstream/"+ref)
            if err != nil {
                return err
            }
        }
    }
    return err
}
