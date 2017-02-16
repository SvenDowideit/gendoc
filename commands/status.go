package commands

import (
	"fmt"
	"os"

	allprojects "github.com/SvenDowideit/gendoc/allprojects"

	"github.com/codegangsta/cli"
)

var logFlag, diffFlag bool

// Status command
var Status = cli.Command{
	Name:  "status",
	Usage: "status versions from " + allprojects.AllProjectsPath + " file",
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
		p := allprojects.GetProjectByName(allprojects.AllProjectsRepo)
		p.Ref = setName
		p.Path = new(string)
		*p.Path = "."
		p.Status(logFlag, diffFlag)

		for _, p := range *projects {
			// TODO: don't ignore errors.
			fmt.Printf("## %s (in %s)\n", p.RepoName, *p.Path)
			p.Status(logFlag, diffFlag)
		}
		return nil
	},
}
