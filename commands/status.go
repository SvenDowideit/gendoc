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
            status(p.RepoName, p.Ref)
        }
        return nil
	},
}

func status(repoPath, ref string) error {
    return allprojects.GitIn(repoPath, "branch", "--contains")
}