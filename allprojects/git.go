package allprojects

import (
	"bufio"
	"os"
	"os/exec"
	"strings"

	"github.com/Sirupsen/logrus"
)

var GithubToken string

func PrintVerboseCommand(cmd *exec.Cmd) {
	logrus.Debugf("executing %q in %s\n", strings.Join(cmd.Args, " "), cmd.Dir)
}

// Execute git commands and output to
// Stdout and Stderr
func Git(args ...string) error {
	cmd := exec.Command("git", args...)
	PrintVerboseCommand(cmd)
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout

	return cmd.Run()
}

func GitIn(dir string, args ...string) error {
	cmd := exec.Command("git", args...)
	cmd.Dir = dir
	PrintVerboseCommand(cmd)
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout

	return cmd.Run()
}

func GitResultsIn(dir string, args ...string) (string, error) {
	cmd := exec.Command("git", args...)
	cmd.Dir = dir
	PrintVerboseCommand(cmd)

	out, err := cmd.Output()
	return string(out), err
}

func GitEnvResultsIn(env []string, dir string, args ...string) (string, error) {
	cmd := exec.Command("git", args...)
	cmd.Dir = dir
	e := os.Environ()
	e = append(e, env...)
	cmd.Env = e
	PrintVerboseCommand(cmd)

	out, err := cmd.Output()
	return string(out), err
}

// GitScannerIn allows the user to parse stdout and stderr - you need to call cmd.Start() and cmd.Wait()
func GitScannerIn(dir string, args ...string) (*bufio.Scanner, *bufio.Scanner, error, *exec.Cmd) {
	cmd := exec.Command("git", args...)
	cmd.Dir = dir
	PrintVerboseCommand(cmd)

	stderr, err := cmd.StderrPipe()
	if err != nil {
		return nil, nil, err, nil
	}
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return nil, nil, err, nil
	}

	return bufio.NewScanner(stdout), bufio.NewScanner(stderr), nil, cmd
}
