package github

import (
	"log"
	"os"
	
	gh "github.com/google/go-github/github"
	"golang.org/x/oauth2"
)

func Render(md string) (html string, err error) {
	client := getClient()
	opt := &gh.MarkdownOptions{Mode: "gfm", Context: "google/go-github"}

	output, _, err := client.Markdown(md, opt)
	if _, ok := err.(*gh.RateLimitError); ok {
		log.Println("EXITING:", err)
		os.Exit(-1)
	}
	return output, err
}

var ghClient *gh.Client

func getClient() *gh.Client {
	if ghClient == nil {
		if gh_token := os.Getenv("GITHUB_TOKEN"); gh_token != "" {
			ts := oauth2.StaticTokenSource(
				&oauth2.Token{AccessToken: gh_token},
			)
			tc := oauth2.NewClient(oauth2.NoContext, ts)
			ghClient = gh.NewClient(tc)
		} else {
			ghClient = gh.NewClient(nil)
		}
	}
	return ghClient
}
