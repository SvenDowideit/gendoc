package allprojects

import (
	"bytes"
    "os"
    "os/exec"
	"fmt"
	"strings"
	"text/template"
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
	fmt.Fprintf(os.Stderr, "executing %q ...\n", strings.Join(cmd.Args, " "))
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

