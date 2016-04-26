package render

import (
	"io/ioutil"
	"log"
	"strings"

	gh "github.com/SvenDowideit/gendoc/render/github"
)

func GithubAPI(markdownFiles []string) {
	for _, file := range markdownFiles {
		log.Println("INFO: ", file)
		input, err := ioutil.ReadFile(file)
		if err != nil {
			log.Println("ERR: ", err)
			continue
		}
		// TODO: remove hugo frontmatter and store for use in the header
		html, err := gh.Render(string(input))
		if err != nil {
			log.Println("ERR: ", err)
			continue
		}
		// TODO: write output to `output` dir, and copy staticFiles too
		outfile := strings.TrimSuffix(file, ".md")+".html"
		if err = ioutil.WriteFile(outfile, []byte(html), 0644); err != nil {
			log.Println("ERR: ", err)
			continue
		}
	}
}
