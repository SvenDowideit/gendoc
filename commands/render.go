package commands

import (
	"fmt"
    "os"
    "os/exec"
    "path/filepath"


	allprojects "github.com/SvenDowideit/gendoc/allprojects"

	"github.com/codegangsta/cli"
)

var Render = cli.Command{
	Name:  "render",
	Usage: "render html of docs checked out.",
	Flags: []cli.Flag{
	},
	Action: func(context *cli.Context) error {
		setName, _, err := allprojects.Load(allprojects.AllProjectsPath)
		if err != nil {
            if os.IsNotExist(err) {
                fmt.Printf("Please run `clone` command first.\n")
            }
			return err
		}
		fmt.Printf("publish-set: %s\n", setName)

        //TODO: confirm that we have the right publish set fetched.
        htmlDir := filepath.Join("../../docs-html/", setName)

        cmd := exec.Command("hugo", "--destination", htmlDir, "--cleanDestinationDir")
        cmd.Dir = filepath.Join("docs-source", setName)

        //PrintVerboseCommand(cmd)
        cmd.Stderr = os.Stderr
        cmd.Stdout = os.Stdout

        return cmd.Run()
	},
}