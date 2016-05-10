package commands

import (
	"fmt"

	// render "github.com/SvenDowideit/gendoc/render"
	allprojects "github.com/SvenDowideit/gendoc/allprojects"

	"github.com/codegangsta/cli"
)

var Clone = cli.Command{
	Name:  "clone",
	Usage: "clone all repos mentioned in the all-projects.yml file",
	Action: func(context *cli.Context) error {

		setName, projects, err := allprojects.Load("./all-projects.yml")
		if err != nil {
			return err
		}
		fmt.Printf("publish-set: %s\n", setName)
		fmt.Printf("projects: %#v\n\n", projects)

		for _, p := range *projects {
			repo, _ := allprojects.GetGitRepo(p)
			fmt.Println(repo)
		}

		return nil
	},
}
