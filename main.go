package main

import (
	"fmt"

	// render "github.com/SvenDowideit/gendoc/render"
	allprojects "github.com/SvenDowideit/gendoc/allprojects"
)



func main() {
	setName, projects, err := allprojects.Load("./all-projects.yml")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("publish-set: %s\n", setName)
	fmt.Printf("projects: %#v\n\n", projects)
	
    for _, p := range *projects {
        repo, _ := allprojects.GetGitRepo(p)
        fmt.Println(repo)
    }
}


