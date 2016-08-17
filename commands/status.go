package commands

import (
	"fmt"
	"os"
	"strings"

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
		status(allprojects.AllProjectsRepo, setName, ".")

		// TODO: confirm what is checked out is the ref from the all-projects
		// TODO: confirm there are no changes - or list them
		for _, p := range *projects {
			// TODO: don't ignore errors.
			fmt.Printf("## %s (in %s)\n", p.RepoName, *p.Path)
			status(p.RepoName, p.Ref, *p.Path)
		}
		return nil
	},
}

func status(repo, ref, path string) {
	// `gendoc update-yaml` :compare
	// the sha we have checked out `git rev-parse --verify --quiet HEAD`
	// the head of the branch `git show-ref upstream/master`
	// and the ref sha in the all-projects
	// if they differ, then tell the user they can update the all-projects
	// add a DOIT flag
	found := false

	// what I plan to show:
	// are we at all-projects ref
	msg, err := hasCheckedOutRef(repo, ref)
	if err == nil {
		fmt.Printf("- %s (as per all-projects.yml)\n", msg)
		found = true
	}
	// are we at upstream/master
	msg, err = hasCheckedOutRef(repo, "master")
	if err == nil {
		fmt.Printf("- %s\n", msg)
		found = true
	}
	// are we based _from_ either of those 2 (ie, new local commits)
	if !found {
		name, err := allprojects.GitResultsIn(repo, "describe", "--abbrev=0")
		name = strings.TrimSpace(name)
		if err == nil {
			full, _ := allprojects.GitResultsIn(repo, "describe")
			full = strings.TrimSpace(full)
			extra := strings.TrimPrefix(full, name)
			if extra == "" {
				fmt.Printf("- checkout is at %s\n", name)
			} else {
				num := strings.Split(extra, "-")
				fmt.Printf("- checkout has %s extra commits after %s\n", num[1], name)
				// TODO: test if `name` == all-projects ref
			}
			// TODO: see if the parent is also master / all-projects
		}
	}
	// is the repo dirty - and what files (ie, `git status`)
	allprojects.GitIn(repo, "status", "-s")
	// logFlag&diffFlag
	if logFlag || diffFlag {
		headSHA, err := getSHA(repo, "HEAD")
		if err != nil {
			fmt.Printf("ERROR: %s\n", err)
			return
		}
		refSHA, err := getSHA(repo, ref)
		if err != nil {
			fmt.Printf("ERROR: (%s) %s\n", ref, err)
			return
		}
		if headSHA != refSHA {
			// TODO: shows all commits and changes, including non docs-dir ones.
			if logFlag {
				allprojects.GitIn(repo, "log", "--oneline", refSHA+".."+headSHA)
			}
			if diffFlag {
				allprojects.GitIn(repo, "diff", refSHA+".."+headSHA)
			}
		}
	}
}
