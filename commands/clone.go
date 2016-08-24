package commands

import (
	"fmt"
	"os"

	allprojects "github.com/docker/gendoc/allprojects"

	"github.com/codegangsta/cli"
)

var cloneSingle, cloneVendoringFlag bool

var Clone = cli.Command{
	Name:  "clone",
	Usage: "clone repos from the " + allprojects.AllProjectsPath + " file",
	Flags: []cli.Flag{
		cli.BoolFlag{
			Name:        "single",
			Usage:       "clone all repositories in " + allprojects.AllProjectsPath + " file",
			Destination: &cloneSingle,
		},
		cli.BoolFlag{
			Name:        "bookkeeping",
			Usage:       "clone docker/docs-html and docker/docs-src too (very large)",
			Destination: &cloneVendoringFlag,
		},
		cli.StringFlag{
			// TODO: consider making this flexible (full git url, org+repo, just org)
			Name:        "origin",
			Usage:       "Set the git origin user/org (for eg 'SvenDowideit')",
			Destination: &originOrgFlag,
		},
	},
	Action: func(context *cli.Context) error {
		setName, projects, err := allprojects.Load(allprojects.AllProjectsPath)
		if os.IsNotExist(err) {
			// clone allProjectsRepo
			// TODO: what if we want londoncalling/docs.docker.com-testing ?
			var project allprojects.Project
			project, err = projects.GetProjectByName(allprojects.AllProjectsRepo)
			if err != nil {
				err = project.CloneRepo(originOrgFlag)
				if err != nil {
					return err
				}
			}
			// try again.
			setName, projects, err = allprojects.Load(allprojects.AllProjectsPath)
		}
		if err != nil {
			return err
		}
		fmt.Printf("publish-set: %s\n", setName)

		if cloneVendoringFlag {
			// get the book keeping repos (ignore the error, its not in the all-projects)
			project, _ := projects.GetProjectByName("docs-source")
			if err := project.CloneRepo(originOrgFlag); err != nil {
				return err
			}
			// get the results repo (ignore the error, its not in the all-projects)
			project, _ = projects.GetProjectByName("docs-html")
			if err := project.CloneRepo(originOrgFlag); err != nil {
				return err
			}
		}

		if !cloneSingle {
			return cloneAll(projects)
		}

		if context.NArg() > 0 {
			name := context.Args()[0]
			project, err := projects.GetProjectByName(name)
			if err != nil {
				return err
			}
			return project.CloneRepo(originOrgFlag)
		}

		return fmt.Errorf("No repository (or --all) specified.")
	},
}

func cloneAll(projects *allprojects.ProjectList) error {
	for _, p := range *projects {
		p.CloneRepo(originOrgFlag)
	}

	return nil
}
