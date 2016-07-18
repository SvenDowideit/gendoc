package allprojects

import (
	"bytes"
	"bufio"
    "os"
    "os/exec"
	"strings"
	"text/template"

	"github.com/Sirupsen/logrus"
)

type Project struct {
	Name     string
	Org      string
	RepoName string `yaml:"repo_name"`
	Ref      string
	Branch   string
	Path     *string
	Target   string
	Ignores  []string
}

func PrintVerboseCommand(cmd *exec.Cmd) {
	logrus.Debugf("executing %q ...\n", strings.Join(cmd.Args, " "))
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

func GitScannerIn(dir string, args ...string) (*bufio.Scanner, *bufio.Scanner, error) {
        cmd := exec.Command("git", args...)
        cmd.Dir = dir
        PrintVerboseCommand(cmd)

	stderr, err := cmd.StderrPipe()
	if err != nil {
	    return nil, nil, err
	}
	stdout, err := cmd.StdoutPipe()
	if err != nil {
	    return nil, nil, err
	}

	return bufio.NewScanner(stdout), bufio.NewScanner(stderr), cmd.Start()
}

func (p Project) GetGitRepo() (string, error) {
	//TODO: extract Template parse
	ghTemplate, err := template.New("repo").Parse("git@github.com:{{.Org}}/{{.RepoName}}")
	var s bytes.Buffer
	if err != nil {
		return "", err
	}
	err = ghTemplate.Execute(&s, p)
	if err != nil {
		return "", err
	}
	return s.String(), nil
}

