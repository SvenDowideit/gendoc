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

		if cloneAllFlag {
			return cloneAll(projects)
		}

		if context.NArg() > 0 {
			name := context.Args()[0]
			project := projects.GetProjectByName(name)
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
	repo, _ := p.GetGitRepo()
	fmt.Println(repo)

    //TODO if it exists, make sure there's a valid remote
    err := allprojects.Git("clone", repo, "--branch", p.Ref, p.RepoName)
	if err != nil {
	    err = allprojects.Git("clone", repo, p.RepoName)
		// TODO checkout?
	}

	return err
}