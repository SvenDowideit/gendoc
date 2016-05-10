package allprojects

import (
	"bytes"
	"fmt"
	"os"
	"strings"
	"text/template"

	"github.com/cloudfoundry-incubator/candiedyaml"
)

type Project struct {
	Name     string
	Org      string
	RepoName string `yaml:"repo_name"`
	Ref      string
	Branch   string
	Path     string
	Target   string
	Ignores  []string
}

type ProjectList []Project

type AllProjects struct {
	PublishSet string      `yaml:"publish-set"`
	Defaults   Project     `yaml:"defaults"`
	Projects   ProjectList `yaml:"projects"`
}

func Load(filename string) (string, *ProjectList, error) {
	file, err := os.Open(filename)
	if err != nil {
		return "", nil, err
	}
	defer file.Close()

	var document AllProjects
	decoder := candiedyaml.NewDecoder(file)
	if err = decoder.Decode(&document); err != nil {
		return "", nil, err
	}

	projects := make(ProjectList, 0)
	for _, p := range document.Projects {
		projects = append(projects, *expand(document.Defaults, p))
	}

	return document.PublishSet, &projects, nil
}

func (projects *ProjectList) GetGitRepos() ([]string, error) {
	repos := make([]string, 0)
	ghTemplate, err := template.New("repo").Parse("git@github.com:{{.Org}}/{{.RepoName}}")
	var s bytes.Buffer
	if err != nil {
		return repos, err
	}
	for _, p := range *projects {
		err = ghTemplate.Execute(&s, p)
		if err != nil {
			return repos, err
		}
		repos = append(repos, s.String())
	}
	return repos, nil
}

func ParseRepoName(name string) Project {
	return Project{
		Org:      "docker",
		Name:     name,
		RepoName: name,
		Ref:      "master",
	}
}

func (p Project) CloneRepo() error {
	repo, _ := p.GetGitRepo()
	fmt.Println(repo)

	return nil
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

func expand(defaults, entry Project) *Project {
	var project Project
	project = entry

	if project.Name == "" {
		project.Name = defaults.Name
	}
	if project.Org == "" {
		project.Org = defaults.Org
	}
	if project.RepoName == "" {
		project.RepoName = defaults.RepoName
	}
	if project.Ref == "" {
		project.Ref = defaults.Ref
	}
	if project.Branch == "" {
		project.Branch = defaults.Branch
	}
	if project.Path == "" {
		project.Path = defaults.Path
	}
	if project.Target == "" {
		project.Target = defaults.Target
	}
	if project.Ignores == nil {
		project.Ignores = defaults.Ignores
	}

	// python template expansions in use atm
	//    repo_name: "{name}"
	//    name: "{repo_name}"
	//    target: "content/{name}"
	if strings.Contains(project.RepoName, "{") {
		project.RepoName = strings.Replace(project.RepoName, "{name}", project.Name, -1)
	}
	if strings.Contains(project.Name, "{") {
		project.Name = strings.Replace(project.Name, "{repo_name}", project.RepoName, -1)
	}
	if strings.Contains(project.Target, "{") {
		project.Target = strings.Replace(project.Target, "{name}", project.Name, -1)
	}

	return &project
}
