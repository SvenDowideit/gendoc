package commands

import (
	"fmt"
	"os"

	allprojects "github.com/SvenDowideit/gendoc/allprojects"

	"github.com/codegangsta/cli"
)

var Release = cli.Command{
	Name:  "release",
	Usage: "Prepare and ship a docs release.",
	Flags: []cli.Flag{
	},
	Subcommands: []cli.Command{
        {
            Name:  "prepare",
            Usage: "Prepare docs release tags and branches.",
            Flags: []cli.Flag{
            },
            Action: func(context *cli.Context) error {
                setName, _, err := allprojects.Load(allprojects.AllProjectsPath)
                if err != nil {
                    return err
                }
                fmt.Printf("publish-set: %s\n", setName)

                return fmt.Errorf("i've got nothin")
            },
        },
        {
            Name:  "push",
            Usage: "Push docs release tags&branches to all the repos.",
            Flags: []cli.Flag{
            },
            Action: func(context *cli.Context) error {
                setName, _, err := allprojects.Load(allprojects.AllProjectsPath)
                if err != nil {
                    return err
                }
                fmt.Printf("publish-set: %s\n", setName)

                return fmt.Errorf("i've got nothin")
            },
        },
        {
            Name:  "release",
            Usage: "Publish docs.",
            Flags: []cli.Flag{
            },
            Action: func(context *cli.Context) error {
                setName, _, err := allprojects.Load(allprojects.AllProjectsPath)
                if err != nil {
                    return err
                }
                fmt.Printf("publish-set: %s\n", setName)

                return fmt.Errorf("i've got nothin")
            },
        },
    },
}

func SOMETHINGCloneRepo(p allprojects.Project) error {
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
