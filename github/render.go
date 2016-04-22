package github

import (
	gh "github.com/google/go-github/github"
)

func Render(md string) (html string, err error) {
	client := getClient()
	opt := &gh.MarkdownOptions{Mode: "gfm", Context: "google/go-github"}

	output, _, err := client.Markdown(md, opt)
	return output, err
}

var ghClient *gh.Client

func getClient() *gh.Client {
	if ghClient == nil {
		ghClient = gh.NewClient(nil)
	}
	return ghClient
}
