package main

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"

	gh "github.com/SvenDowideit/gendoc/github"
)

func main() {
	markdownFiles := []string{}
	err := filepath.Walk("./docs", func(path string, f os.FileInfo, err error) error {
		if err != nil {
			log.Println("ERR: ", err)
		}
		if ! f.IsDir() {
			if strings.HasSuffix(path, ".md") {
				markdownFiles = append(markdownFiles, path)
			}
		}
		return nil
	})
	if err != nil {
		log.Println("ERR: ", err)
	}
	
	for _, file := range markdownFiles {
		log.Println("INFO: ", file)
		input, err := ioutil.ReadFile(file)
		if err != nil {
			log.Println("ERR: ", err)
			continue
		}
		html, err := gh.Render(string(input))
		if err != nil {
			log.Println("ERR: ", err)
			continue
		}
		outfile := strings.TrimSuffix(file, ".md")+".html"
		if err = ioutil.WriteFile(outfile, []byte(html), 0644); err != nil {
			log.Println("ERR: ", err)
			continue
		}
	}
	
}
