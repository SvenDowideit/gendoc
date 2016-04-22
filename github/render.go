package github

import (
	"fmt"

	"github.com/google/go-github/github"
)

func Render() {
	client := github.NewClient(nil)

	input := "# heading\n\n1. Link to issue #1\n"
	opt := &github.MarkdownOptions{Mode: "gfm", Context: "google/go-github"}

	output, _, err := client.Markdown(input, opt)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(output)
}
