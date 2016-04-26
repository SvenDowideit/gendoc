package main

import (
	"log"
	"os"
	"path/filepath"
	"strings"

	render "github.com/SvenDowideit/gendoc/render"
)

type SiteData struct {
	MarkdownFiles []string
	StaticFiles   []string
}

func main() {
	site := SiteData{
		MarkdownFiles: []string{},
		StaticFiles:   []string{},
	}

	gatherFilenames("./docs", &site)

	render.GithubAPI("./output_gh", site.MarkdownFiles)
	render.CopyStaticFiles("./output_gh", site.StaticFiles)
}

func gatherFilenames(docsDir string, site *SiteData) {
	err := filepath.Walk(docsDir, func(path string, f os.FileInfo, err error) error {
		if err != nil {
			log.Println("ERR: ", err)
		}
		if !f.IsDir() {
			if strings.HasSuffix(path, ".md") {
				site.MarkdownFiles = append(site.MarkdownFiles, path)
			} else {
				site.StaticFiles = append(site.StaticFiles, path)
			}
		}
		return nil
	})
	if err != nil {
		log.Println("ERR: ", err)
	}
}
