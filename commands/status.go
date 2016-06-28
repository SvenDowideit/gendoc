package commands

import (
	"fmt"
	"os"

	allprojects "github.com/SvenDowideit/gendoc/allprojects"

	"github.com/codegangsta/cli"
)

var Status = cli.Command{
	Name:  "status",
	Usage: "status versions from "+allprojects.AllProjectsPath+" file",
	Flags: []cli.Flag{
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
            fmt.Printf("-- %s\n", p.RepoName)
            status(p.RepoName)

            // `gendoc update-yaml` :compare 
            // the sha we have checked out `git rev-parse --verify --quiet HEAD`
            // the head of the branch `git show-ref origin/master`
            // and the ref sha in the all-projects
            // if they differ, then tell the user they can update the all-projects
            // add a DOIT flag
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