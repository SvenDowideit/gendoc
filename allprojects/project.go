package allprojects

import (
	"bytes"
	"bufio"
	"net/http"
	"os"
	"os/exec"
	"strings"
	"text/template"

	"golang.org/x/oauth2"
	"github.com/Sirupsen/logrus"
	"github.com/google/go-github/github"
)

var GithubToken string

type Project struct {
	Name     string
	Org      string
	RepoName string `yaml:"repo_name"`
	Ref      string
	Branch   string
	Path     *string
	Target   string
	Ignores  []string
	Version  string
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


func GetPRInfo(org, repo string, pr int) (lables, milstone string, err error) {
	var tc *http.Client
	if GithubToken != "" {
		logrus.Debugf("using GitHub API token")
		ts := oauth2.StaticTokenSource(
			&oauth2.Token{AccessToken: GithubToken},
		)
		tc = oauth2.NewClient(oauth2.NoContext, ts)
	}
	// TODO:use token env for oauth
	client := github.NewClient(tc)
	issue, _, err := client.Issues.Get(org, repo, pr)
	if err != nil {
		return "", "", err
	}
	labels := ""
	if issue.Labels != nil {
		for _, l := range issue.Labels {
			labels += *l.Name + " "
		}
	}
	milestone := ""
	if issue.Milestone != nil {
		milestone = *issue.Milestone.Title
	}
	return labels, milestone, err
}
