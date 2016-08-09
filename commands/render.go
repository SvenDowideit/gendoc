package commands

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"

	allprojects "github.com/SvenDowideit/gendoc/allprojects"

	"github.com/codegangsta/cli"
)

var vendorFlag, diskFlag, doitFlag bool
var portFlag int

var Render = cli.Command{
	Name:  "render",
	Usage: "render html of docs checked out.",
	Flags: []cli.Flag{
		cli.BoolTFlag{
			Name:        "vendor",
			Usage:       "vendor changes into docs-source (disable to ignore new changes)",
			Destination: &vendorFlag,
		},
		cli.BoolFlag{
			Name:        "disk",
			Usage:       "Render html to the `docs-html/vMAJOR.MINOR/` dir - no webserver",
			Destination: &diskFlag,
		},
		cli.IntFlag{
			Name:        "port",
			Usage:       "TCP port the rendered docs should be served from",
			Destination: &portFlag,
			Value:       8000,
		},
		cli.BoolFlag{
			Name:        "doit",
			Usage:       "render even though some projects are missing",
			Destination: &doitFlag,
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
		if vendorFlag {
			err = VendorSource(setName, projects)
			if err != nil {
				return err
			}
		}

		//TODO: confirm that we have the right publish set fetched.
		htmlDir := filepath.Join("../../docs-html/", setName)

		// TODO --watch won't work - need to also watch the repo dirs and fetch in background
		// TODO what about the --baseUrl
		opts := []string{
			"--destination", htmlDir,
			"--cleanDestinationDir",
		}
		var cmd *exec.Cmd
		if diskFlag {
			// TODO: find a way to tell hugo not to insert the live-reload html
		} else {
			hugoCmd := []string{"serve"}

			opts = append(hugoCmd, opts...)
			opts = append(opts,
				"--renderToDisk",
				"--port", fmt.Sprintf("%d", portFlag),
				"--bind", "0.0.0.0",
				"--watch")
		}
		cmd = exec.Command("hugo", opts...)

		cmd.Dir = filepath.Join("docs-source", setName)

		//PrintVerboseCommand(cmd)
		cmd.Stderr = os.Stderr
		cmd.Stdout = os.Stdout

		return cmd.Run()
	},
}

func VendorSource(setName string, projects *allprojects.ProjectList) error {
	// TODO add a watch at the end.
	for _, p := range *projects {
		if p.RepoName == "docs.docker.com" {
			continue
		}
		//TODO exclude?
		from := filepath.Join(p.RepoName, *p.Path)
		to := filepath.Join("docs-source", setName, p.Target)
		os.MkdirAll(to, 0755)
		fmt.Printf("copy %s TO %s\n", from, to)
		err := CopyDir(from, to)
		if !doitFlag && err != nil {
			fmt.Println("HINT: add the --doit flag to render when repos are missing")
			return err
		}
		// TODO: write build.json file
	}
	return nil
}

// Copies file source to destination dest.
func CopyFile(source string, dest string) (err error) {
	//fmt.Printf("CopyFile %s TO %s\n", source, dest)

	sf, err := os.Open(source)
	if err != nil {
		return err
	}
	defer sf.Close()
	df, err := os.Create(dest)
	if err != nil {
		return err
	}
	defer df.Close()
	_, err = io.Copy(df, sf)
	if err == nil {
		si, err := os.Stat(source)
		if err != nil {
			err = os.Chmod(dest, si.Mode())
		}

	}

	return
}

// Recursively copies a directory tree, attempting to preserve permissions.
// Source directory must exist, destination directory must *not* exist.
func CopyDir(source string, dest string) (err error) {

	// get properties of source dir
	fi, err := os.Stat(source)
	if err != nil {
		return err
	}

	if !fi.IsDir() {
		return &CustomError{"Source is not a directory"}
	}

	// create dest dir

	err = os.MkdirAll(dest, fi.Mode())
	if err != nil {
		return err
	}

	entries, err := ioutil.ReadDir(source)

	for _, entry := range entries {

		sfp := source + "/" + entry.Name()
		dfp := dest + "/" + entry.Name()
		if entry.IsDir() {
			err = CopyDir(sfp, dfp)
			if err != nil {
				log.Println(err)
			}
		} else {
			// perform copy
			err = CopyFile(sfp, dfp)
			if err != nil {
				log.Println(err)
			}
		}

	}
	return
}

// A struct for returning custom error messages
type CustomError struct {
	What string
}

// Returns the error message defined in What as a string
func (e *CustomError) Error() string {
	return e.What
}
