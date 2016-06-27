package commands

import (
	"fmt"
	"os"
    "io"
    "io/ioutil"
    "log"
    "path/filepath"

	allprojects "github.com/SvenDowideit/gendoc/allprojects"

	"github.com/codegangsta/cli"
)

var Fetch = cli.Command{
	Name:  "fetch",
	Usage: "fetch versions from "+allprojects.AllProjectsPath+" file into publish-set dir",
	Flags: []cli.Flag{
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

        // TODO add a watch at the end.
        for _, p := range *projects {
            //TODO exclude?
            from := filepath.Join(p.RepoName, p.Path)
            to := filepath.Join(setName, p.Target)
            os.MkdirAll(to, 0755)
	        fmt.Printf("copy %s TO %s\n", from, to)
            err = CopyDir(from, to)
            if err != nil {
                return err
            }
        }
        return nil
	},
}

// Copies file source to destination dest.
func CopyFile(source string, dest string) (err error) {
	fmt.Printf("CopyFile %s TO %s\n", source, dest)

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
