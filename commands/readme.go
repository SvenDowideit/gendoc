package commands

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"strings"

	"github.com/codegangsta/cli"
)

var globalTestDir string

var Readme = cli.Command{
	Name:  "readme",
	Usage: "Parse the README file and update any inline command examples",
	Flags: []cli.Flag{},
	Action: func(context *cli.Context) error {
		var err error
		globalTestDir, err = ioutil.TempDir("", "example")
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("Using %s to run commands\n", globalTestDir)
		defer os.RemoveAll(globalTestDir) // clean up

		// read in the README.md
		inFile, _ := os.Open("README.md")
		defer inFile.Close()
		scanner := bufio.NewScanner(inFile)
		scanner.Split(bufio.ScanLines)

		inCode := false // assume we start a file not in a code section
		for scanner.Scan() {
			line := scanner.Text()

			// TODO: deal with trailing chars, and leading whitespace
			if line == "```" {
				fmt.Println(line)
				inCode = !inCode
			} else {
				if inCode {
					if strings.HasPrefix(line, "$") {
						fmt.Println(line)
						cmd := strings.TrimSpace(strings.TrimPrefix(line, "$"))
						// run it and print stdout&stderr
						// TODO check for trailing \
						// and deal with &&
						out, _ := runCmd(cmd)
						fmt.Println(out)
					}
				} else {
					fmt.Println(line)
				}
			}
		}

		return nil
	},
}

func runCmd(command string) (string, error) {
	args := strings.Split(command, " ")
	cmd := exec.Command(args[0], args[1:]...)
	cmd.Dir = globalTestDir
	//	PrintVerboseCommand(cmd)

	// TODO: need a timeout for the hugo service
	out, err := cmd.Output()
	return string(out), err
}
