package main

import (
	"fmt"

	gh "github.com/SvenDowideit/gendoc/github"
)

func main() {
	input := "# heading\n\n1. Link to issue #1\n"
	html, err := gh.Render(input)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(html)
}
