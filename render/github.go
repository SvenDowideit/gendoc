package render

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"

	gh "github.com/SvenDowideit/gendoc/render/github"
)

func GithubAPI(outputDir string, markdownFiles []string) {
	for _, file := range markdownFiles {
		log.Println("INFO: <<", file)
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
		outfile := filepath.FromSlash(filepath.Join(outputDir, strings.TrimSuffix(file, ".md")+".html"))
		if err := os.MkdirAll(filepath.Dir(outfile), 0755); err != nil {
			log.Println("ERR: ", err)
			continue
		}
		if err = ioutil.WriteFile(outfile, []byte(html), 0644); err != nil {
			log.Println("ERR: ", err)
			continue
		}
		log.Println("INFO: >>", outfile)
	}
}
