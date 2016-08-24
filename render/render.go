package render

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"

	gh "github.com/docker/gendoc/render/github"
	mmark "github.com/docker/gendoc/render/mmark"
)

type SiteData struct {
	MarkdownFiles []string
	StaticFiles   []string
}

// RenderDocsDir will find all the markdown files, and static files in the docs subdir
// and convert them to html in the output_mmark dir
func RenderDocsDir() {
	site := SiteData{
		MarkdownFiles: []string{},
		StaticFiles:   []string{},
	}

	gatherFilenames("./docs", &site)

	//	GithubAPI("./output_gh", site.MarkdownFiles)
	//	CopyStaticFiles("./output_gh", site.StaticFiles)

	MMark("./output_mmark", site.MarkdownFiles)
	CopyStaticFiles("./output_mmark", site.StaticFiles)
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

type renderFunc func(string) (string, error)

func GithubAPI(outputDir string, markdownFiles []string) {
	processFiles(outputDir, markdownFiles, gh.Render)
}
func MMark(outputDir string, markdownFiles []string) {
	processFiles(outputDir, markdownFiles, mmark.Render)
}

func processFiles(outputDir string, markdownFiles []string, render renderFunc) {
	for _, file := range markdownFiles {
		processFile(outputDir, file, render)
	}
}

func processFile(outputDir, file string, render renderFunc) error {
	log.Println("INFO: <<", file)
	input, err := ioutil.ReadFile(file)
	if err != nil {
		log.Println("ERR: ", err)
		return err
	}
	// TODO: remove hugo frontmatter and store for use in the header
	html, err := render(string(input))
	if err != nil {
		log.Println("ERR: ", err)
		return err
	}
	outfile := filepath.FromSlash(filepath.Join(outputDir, strings.TrimSuffix(file, ".md")+".html"))
	if err := os.MkdirAll(filepath.Dir(outfile), 0755); err != nil {
		log.Println("ERR: ", err)
		return err
	}
	if err = ioutil.WriteFile(outfile, []byte(html), 0644); err != nil {
		log.Println("ERR: ", err)
		return err
	}
	log.Println("INFO: >>", outfile)
	return nil
}
