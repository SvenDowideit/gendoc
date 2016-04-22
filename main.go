package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	gh "github.com/SvenDowideit/gendoc/github"
)

func main() {
	fileList := []string{}
	err := filepath.Walk("./docs", func(path string, f os.FileInfo, err error) error {
		if err != nil {
			fmt.Println("ERR: ", err)
		}
		fileList = append(fileList, path)
		return nil
	})
	if err != nil {
		fmt.Println("ERR: ", err)
	}
	
	for _, file := range fileList {
		fmt.Println("INFO: ", file)
		if !strings.HasSuffix(file, ".md") {
			fmt.Println("INFO: skipping non markdown file")
			continue
		}
		input, err := ioutil.ReadFile(file)
		if err != nil {
			fmt.Println("ERR: ", err)
			continue
		}
		html, err := gh.Render(string(input))
		if err != nil {
			fmt.Println("ERR: ", err)
			continue
		}
		outfile := strings.TrimSuffix(file, ".md")+".html"
		if err = ioutil.WriteFile(outfile, []byte(html), 0644); err != nil {
			fmt.Println("ERR: ", err)
			continue
		}
	}
	
}
