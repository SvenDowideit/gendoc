package allprojects

import (
	"bytes"
	"fmt"
	"net/http"
	"os"
	"strings"
	"text/template"

	"github.com/Sirupsen/logrus"
	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
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
	Version  string
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

// AddRemote adds a git remote repo
func (p Project) AddRemote(name, org string) error {
	p.Org = org
	repo, _ := p.GetGitRepo()
	return GitIn(p.RepoName, "remote", "add", name, repo)
}

//TODO: bail out if there are local commits, or isdirty
func (p Project) Checkout(fetchFlag, resetFlag bool, originOrgFlag string) error {
	if fetchFlag {
		err := GitIn(p.RepoName, "fetch", "--all")
		if err != nil {
			return err
		}
	}

	// exit happy if the sha of HEAD == the SHA that the ref points to (not the sha of the tag)
	msg, err := p.HasCheckedOutRef(p.Ref)
	if err == nil {
		fmt.Printf("Same as all-projects.yml: %s\n", msg)
		return nil
	}

	err = GitIn(p.RepoName, "checkout", p.Ref)
	if err != nil {
		// do a fetch, in case it exists in remote
		err = GitIn(p.RepoName, "fetch", "upstream", p.Ref+":remotes/upstream/"+p.Ref)
		if err != nil {
			// Last resourt, fetch all upstream, and undo depth
			err = GitIn(p.RepoName, "fetch", "--all")
			if err != nil {
				return err
			}
			err = GitIn(p.RepoName, "fetch", "--tag", "upstream")
			if err != nil {
				return err
			}
		}
		err = GitIn(p.RepoName, "checkout", p.Ref)
		if err != nil {
			err = GitIn(p.RepoName, "checkout", "-b", p.Ref, "remotes/upstream/"+p.Ref)
			if err != nil {
				return err
			}
		}
	}
	if resetFlag {
		if _, err := GitResultsIn(p.RepoName, "show-ref", "--hash", "upstream/"+p.Ref); err == nil {
			// its not a SHA, so we can reset
			err = GitIn(p.RepoName, "reset", "--hard", "upstream/"+p.Ref)
			if err != nil {
				return err
			}
		}
	}
	if originOrgFlag != "" {
		p.AddRemote("origin", originOrgFlag)
	}
	return err
}

// GetSHA returns the SHA of `ref` in this repo
func (p Project) GetSHA(ref string) (sha string, err error) {
	sha, err = GitResultsIn(p.RepoName, "log", "-1", "--format=%H", ref)
	return strings.TrimSpace(sha), err
}

// HasCheckedOutRef returns a message and nil error if the repo has `ref` checked out
func (p Project) HasCheckedOutRef(ref string) (string, error) {
	// exit happy if the sha of HEAD == the SHA that the ref points to (not the sha of the tag)
	headSHA, err := p.GetSHA("HEAD")
	if err != nil {
		// if we can't get the SHA of HEAD, we're dead
		return "", err
	}
	logrus.Debugf("compare (%s) to (%s)\n", headSHA, ref)
	if headSHA == ref {
		// the all-projects ref is a SHA
		return fmt.Sprintf("your checkout %s is at %s", headSHA, ref), nil
	}
	// is ref an upstream branch?
	if refSHA, err := p.GetSHA("upstream/" + ref); err == nil {
		logrus.Debugf("compare (%s) to (%s)\n", headSHA, refSHA)
		// if we got that ok, we don't need a checkout / fetch
		if headSHA == refSHA {
			return fmt.Sprintf("your checkout %s is at %s", headSHA, "upstream/"+ref), nil
		}
	}
	// is ref a tag?
	if tagSHA, err := GitResultsIn(p.RepoName, "show-ref", "--hash", "refs/tags/"+ref); err == nil {
		tagSHA = strings.TrimSpace(tagSHA)
		if refSHA, err := p.GetSHA(tagSHA); err == nil {
			logrus.Debugf("compare (%s) to (%s)\n", headSHA, refSHA)
			// if we got that ok, we don't need a checkout / fetch
			if headSHA == refSHA {
				return fmt.Sprintf("your checkout %s is at %s", headSHA, "refs/tags/"+ref), nil
			}
		}
	}
	return "", fmt.Errorf("Error: repository not at %s", ref)
}

func (p Project) Status(logFlag, diffFlag bool) {
	// `gendoc update-yaml` :compare
	// the sha we have checked out `git rev-parse --verify --quiet HEAD`
	// the head of the branch `git show-ref upstream/master`
	// and the ref sha in the all-projects
	// if they differ, then tell the user they can update the all-projects
	// add a DOIT flag
	found := false

	// what I plan to show:
	// are we at all-projects ref
	msg, err := p.HasCheckedOutRef(p.Ref)
	if err == nil {
		fmt.Printf("- %s (as per all-projects.yml)\n", msg)
		found = true
	}
	// are we at upstream/master
	msg, err = p.HasCheckedOutRef("master")
	if err == nil {
		fmt.Printf("- %s\n", msg)
		found = true
	}
	// are we based _from_ either of those 2 (ie, new local commits)
	if !found {
		name, err := GitResultsIn(p.RepoName, "describe", "--abbrev=0")
		name = strings.TrimSpace(name)
		if err == nil {
			full, _ := GitResultsIn(p.RepoName, "describe")
			full = strings.TrimSpace(full)
			extra := strings.TrimPrefix(full, name)
			if extra == "" {
				fmt.Printf("- checkout is at %s\n", name)
			} else {
				num := strings.Split(extra, "-")
				fmt.Printf("- checkout has %s extra commits after %s\n", num[1], name)
				// TODO: test if `name` == all-projects ref
			}
			// TODO: see if the parent is also master / all-projects
		}
	}
	// is the repo dirty - and what files (ie, `git status`)
	GitIn(p.RepoName, "status", "-s")
	// logFlag&diffFlag
	if logFlag || diffFlag {
		headSHA, err := p.GetSHA("HEAD")
		if err != nil {
			fmt.Printf("ERROR: %s\n", err)
			return
		}
		refSHA, err := p.GetSHA(p.Ref)
		if err != nil {
			fmt.Printf("ERROR: (%s) %s\n", p.Ref, err)
			return
		}
		if headSHA != refSHA {
			// TODO: shows all commits and changes, including non docs-dir ones.
			if logFlag {
				GitIn(p.RepoName, "log", "--oneline", refSHA+".."+headSHA)
			}
			if diffFlag {
				GitIn(p.RepoName, "diff", refSHA+".."+headSHA)
			}
		}
	}
}

func (p Project) CloneRepo(originOrgFlag string) error {
	fmt.Printf("-- %s\n", p.Name)
	//TODO if it exists, make sure there's a valid remote
	if _, err := os.Stat(p.RepoName); os.IsNotExist(err) {
		repo, _ := p.GetGitRepo()
		fmt.Printf("Cloning from %s\n", repo)
		err = Git("clone", "--origin", "upstream", repo, p.RepoName)
		if originOrgFlag != "" {
			p.AddRemote("origin", originOrgFlag)
		}
		return err
	}

	fmt.Printf("Dir already exists\n")
	return nil
}
