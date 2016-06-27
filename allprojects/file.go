package allprojects

import (
	"bytes"
	"os"
	"strings"
	"text/template"

	"github.com/cloudfoundry-incubator/candiedyaml"
)

var AllProjectsRepo = "docs.docker.com"
var AllProjectsPath = "./docs.docker.com/all-projects.yml"

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
		projects = append(projects, *expandDefaults(document.Defaults, p))
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

func (projects *ProjectList) GetProjectByName(name string) Project {

	// TODO: I presume this is naughty :)
	if projects != nil {
		for _, p := range *projects {
			if p.Name == name {
				return p
			}
		}
	}
	// project not in all-projects.yml
	return makeProject(name)
}
func makeProject(name string) Project {
	return Project{
		Org:      "docker",
		Name:     name,
		RepoName: name,
		Ref:      "master",
	}
}

// expandDefaults is only used to default values when parsing the yaml
func expandDefaults(defaults, entry Project) *Project {
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
	if project.Path == nil {
		project.Path = new(string)
		*project.Path = *defaults.Path
	} else {
		if *project.Path == "" {
			// yeah, !!null - thanks.
			*project.Path = "."
		}
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
