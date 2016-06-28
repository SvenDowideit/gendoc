package commands

import (
	"fmt"
    "os"
    "os/exec"
    "path/filepath"


	allprojects "github.com/SvenDowideit/gendoc/allprojects"

	"github.com/codegangsta/cli"
)

var fetchFlag bool

var Serve = cli.Command{
	Name:  "serve",
	Usage: "serve html of docs checked out.",
	Flags: []cli.Flag{
		cli.BoolTFlag{
			Name:        "fetch",
			Usage:       "do a fetch of files from the checked out repos first",
			Destination: &fetchFlag,
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

        if fetchFlag {
            err = DoFetch(setName, projects)
			if err != nil {
				return err
			}
        }

        //TODO: confirm that we have the right publish set fetched.
        htmlDir := filepath.Join("../../docs-html/", setName)

        // TODO --watch won't work - need to also watch the repo dirs and fetch in background
        cmd := exec.Command("hugo", "serve", 
            "--renderToDisk",
            "--destination", htmlDir,
            "--port", "8080",
            "--cleanDestinationDir",
            "--watch")
        cmd.Dir = filepath.Join("docs-source", setName)

        //PrintVerboseCommand(cmd)
        cmd.Stderr = os.Stderr
        cmd.Stdout = os.Stdout

        return cmd.Run()
	},
}