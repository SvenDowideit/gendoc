package commands

import (
	"fmt"

	allprojects "github.com/SvenDowideit/gendoc/allprojects"

	"github.com/codegangsta/cli"
)

var cloneAllFlag bool

var Clone = cli.Command{
	Name:  "clone",
	Usage: "clone all repos mentioned in the all-projects.yml file",
	Flags: []cli.Flag{
		cli.BoolFlag{
			Name:        "all",
			Usage:       "clone all repositories in all-projects.yml file",
			Destination: &cloneAllFlag,
		},
	},
	Action: func(context *cli.Context) error {
		if cloneAllFlag {
			return cloneAll()
		}

		if context.NArg() > 0 {
			name := context.Args()[0]
			project := allprojects.ParseRepoName(name)
			project.CloneRepo()
		}

		return nil
	},
}

func cloneAll() error {
	setName, projects, err := allprojects.Load("./all-projects.yml")
	if err != nil {
		return err
	}
	fmt.Printf("publish-set: %s\n", setName)
	// fmt.Printf("projects: %#v\n\n", projects)

	for _, p := range *projects {
		p.CloneRepo()
	}

	return nil
}
