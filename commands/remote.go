package commands

import (
	"fmt"
	"os"

	allprojects "github.com/SvenDowideit/gendoc/allprojects"

	"github.com/codegangsta/cli"
)

var Remote = cli.Command{
	Name:  "remote",
	Usage: "Add a git remote",
	Flags: []cli.Flag{
	},
	Action: func(context *cli.Context) error {
		if context.NArg() != 2 {
            return fmt.Errorf("Please specify the remote name and repo org to add")
		}
		name := context.Args()[0]
		org := context.Args()[1]

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
            addRemote(p, name, org)
        }
        return nil
	},
}

func addRemote(p allprojects.Project, name, org string) error {
    p.Org = org
   	repo, _ := p.GetGitRepo()
    return allprojects.GitIn(p.RepoName, "remote", "add", name, repo)
}