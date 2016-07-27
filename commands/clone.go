package commands

import (
	"fmt"
	"os"

	allprojects "github.com/SvenDowideit/gendoc/allprojects"

	"github.com/codegangsta/cli"
)

var cloneSingle, cloneVendoringFlag bool

var Clone = cli.Command{
	Name:  "clone",
	Usage: "clone repos from the "+allprojects.AllProjectsPath+" file",
	Flags: []cli.Flag{
		cli.BoolFlag{
			Name:        "single",
			Usage:       "clone all repositories in "+allprojects.AllProjectsPath+" file",
			Destination: &cloneSingle,
		},
		cli.BoolFlag{
			Name:        "bookkeeping",
			Usage:       "clone docker/docs-html and docker/docs-src too (very large)",
			Destination: &cloneVendoringFlag,
		},
		//TODO: add an shallow clone flag
	},
	Action: func(context *cli.Context) error {
		setName, projects, err := allprojects.Load(allprojects.AllProjectsPath)
		if os.IsNotExist(err) {
			// clone allProjectsRepo
			// TODO: what if we want londoncalling/docs.docker.com-testing ?
			project, err := projects.GetProjectByName(allprojects.AllProjectsRepo)
			if err != nil {
				return err
			}
			err = CloneRepo(project)
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

		if cloneVendoringFlag {
			// get the book keeping repos
			project, err := projects.GetProjectByName("docs-source")
			if err != nil {
				return err
			}
			if err := CloneRepo(project); err != nil {
                                return err
                        }
			project, err = projects.GetProjectByName("docs-html")
			if err != nil {
				return err
			}
			if err := CloneRepo(project); err != nil {
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
			return CloneRepo(project)
		}

		return fmt.Errorf("No repository (or --all) specified.")
	},
}

func cloneAll(projects *allprojects.ProjectList) error {
	for _, p := range *projects {
		CloneRepo(p)
	}

	return nil
}

func CloneRepo(p allprojects.Project) error {
	fmt.Printf("-- %s\n", p.Name)
	//TODO if it exists, make sure there's a valid remote
	if _, err := os.Stat(p.RepoName); os.IsNotExist(err) {
		repo, _ := p.GetGitRepo()
		fmt.Printf("Cloning from %s\n", repo)

		//err := allprojects.Git("clone", repo, "--branch", p.Ref, p.RepoName)
		//if err != nil {
			err = allprojects.Git("clone", "--origin", "upstream", repo, p.RepoName)
		//}
		return err
	} else {
		fmt.Printf("Dir already exists\n")
	}
	return nil
}
