package commands

import (
	"fmt"
	"os"

	allprojects "github.com/SvenDowideit/gendoc/allprojects"

	"github.com/codegangsta/cli"
)

var cloneAllFlag = true

var Clone = cli.Command{
	Name:  "clone",
	Usage: "clone repos from the "+allprojects.AllProjectsPath+" file",
	Flags: []cli.Flag{
		cli.BoolFlag{
			Name:        "all",
			Usage:       "clone all repositories in "+allprojects.AllProjectsPath+" file",
			Destination: &cloneAllFlag,
		},
		//TODO: add an shallow clone flag
	},
	Action: func(context *cli.Context) error {
		setName, projects, err := allprojects.Load(allprojects.AllProjectsPath)
		if os.IsNotExist(err) {
			// clone allProjectsRepo
			// TODO: what if we want londoncalling/docs.docker.com-testing ?
			project := projects.GetProjectByName(allprojects.AllProjectsRepo)
			err = project.CloneRepo()
			if err != nil {
				return err
			}
			// try again.
			setName, projects, err = allprojects.Load(allprojects.AllProjectsPath)
		}
		if err != nil {
			return err
		}
		fmt.Printf("publish-set: %s\n", setName)

		if cloneAllFlag {
			return cloneAll(projects)
		}

		if context.NArg() > 0 {
			name := context.Args()[0]
			project := projects.GetProjectByName(name)
			return project.CloneRepo()
		}

		return fmt.Errorf("No repository (or --all) specified.")
	},
}

func cloneAll(projects *allprojects.ProjectList) error {
	for _, p := range *projects {
		p.CloneRepo()
	}

	return nil
}
