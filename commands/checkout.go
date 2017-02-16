package commands

import (
	"fmt"
	"os"

	"github.com/SvenDowideit/gendoc/allprojects"

	"github.com/codegangsta/cli"
)

var fetchFlag bool
var resetFlag bool
var originOrgFlag string

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
		cli.StringFlag{
			// TODO: consider making this flexible (full git url, org+repo, just org)
			Name:  "origin",
			Usage: "Set the git origin user/org (for eg 'SvenDowideit')",
			// Value:       "upstream/master",
			Destination: &originOrgFlag,
		},
	},
	Action: func(context *cli.Context) error {
		// TODO: checkout what's in the current file - we might be testing a branch
		if context.NArg() == 1 {
			publishSetBranch := context.Args()[0]
			fmt.Printf("Checking out %s %s.\n", allprojects.AllProjectsRepo, publishSetBranch)
			p := allprojects.GetProjectByName(allprojects.AllProjectsRepo)
			p.Ref = publishSetBranch
			err := p.Checkout(fetchFlag, resetFlag, originOrgFlag)
			if err != nil {
				return err
			}
		} else {
			fmt.Printf("Using the docs.docker.com/all-projects.yml as is.\n")

		}

		//TODO need to fetch&reset docs-html and docs-src

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
			p.Checkout(fetchFlag, resetFlag, originOrgFlag)
		}
		return nil
	},
}
