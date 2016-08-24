package commands

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	allprojects "github.com/SvenDowideit/gendoc/allprojects"
	"github.com/docker/markdownlint/checkers"
	"github.com/docker/markdownlint/data"
	"github.com/docker/markdownlint/linereader"

	"github.com/codegangsta/cli"
)

var externalLinksFlag bool

// Test command
var Test = cli.Command{
	Name:  "test",
	Usage: "Run the markdown checker",
	Flags: []cli.Flag{
		cli.BoolTFlag{
			Name:        "vendor",
			Usage:       "vendor changes into docs-source (disable to ignore new changes)",
			Destination: &vendorFlag,
		},
		cli.BoolFlag{
			Name:        "external",
			Usage:       "test external URLS",
			Destination: &externalLinksFlag,
		},
		cli.BoolFlag{
			Name:        "noisy",
			Usage:       "tell me about all skipped PR's",
			Destination: &noisyFlag,
		},
	},
	Action: func(context *cli.Context) error {
		setName, projects, err := allprojects.Load(allprojects.AllProjectsPath)
		if err != nil {
			if os.IsNotExist(err) {
				fmt.Printf("Please run `clone` command first.\n")
			}
			return err
		}
		fmt.Printf("publish-set: %s\n", setName)
		if vendorFlag {
			err = VendorSource(setName, projects)
			if err != nil {
				return err
			}
		}
		docsSourceDir := filepath.Join("docs-source", setName, "content")
		markdownLint(docsSourceDir, "")

		return nil
	},
}

// this is a slightly modified copy of main() in the markdownlint repo.
func markdownLint(dir, filter string) {
	data.AllFiles = make(map[string]*data.FileDetails)

	fmt.Printf("\nFinding markdown files in %s\n", dir)
	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			data.ErrorLog("%s\n", err)
			return err
		}
		data.VerboseLog("FOUND: %s\n", path)
		if info.IsDir() {
			return nil
		}
		file, err := filepath.Rel(dir, path)
		if err != nil {
			data.ErrorLog("%s\n", err)
			return err
		}
		// verboseLog("\t walked to %s\n", file)
		data.AddFile(file, path)
		return nil
	})
	if err != nil {
		data.ErrorLog("%s\n", err)
		os.Exit(-1)
	}

	count := 0
	for file, details := range data.AllFiles {
		if !strings.HasPrefix(file, filter) {
			data.VerboseLog("FILTERED: %s\n", file)
			continue
		}
		if !strings.HasSuffix(file, ".md") {
			data.VerboseLog("SKIPPING: %s\n", file)
			continue
		}
		// fmt.Printf("opening: %s\n", file)
		count++
		if count%100 == 0 {
			fmt.Printf("\topened %d files so far\n", count)
		}

		reader, err := linereader.OpenReader(details.FullPath)
		if err != nil {
			data.ErrorLog("%s\n", err)
			data.AllFiles[file].FormatErrorCount++
		}

		err = checkers.CheckHugoFrontmatter(reader, file)
		if err != nil {
			data.ErrorLog("(%s) frontmatter: %s\n", file, err)
		}

		if draft, ok := data.AllFiles[file].Meta["draft"]; ok || draft == "true" {
			data.VerboseLog("Draft=%s: SKIPPING %s link check.\n", draft, file)
		} else {
			//fmt.Printf("Draft=%s: %s link check.\n", draft, file)
			err = checkers.CheckMarkdownLinks(reader, file)
			if err != nil {
				// this only errors if there is a fatal issue
				data.ErrorLog("(%s) links: %s\n", file, err)
				data.AllFiles[file].FormatErrorCount++
			}
		}
		reader.Close()
	}

	fmt.Printf("Starting to test links (Filter = %s)\n", filter)
	checkers.TestLinks(filter, externalLinksFlag)

	// TODO (JIRA: DOCS-181): Title, unique across products if not, file should include an {identifier}

	summaryFileName := "markdownlint.summary.txt"
	f, err := os.Create(summaryFileName)
	if err == nil {
		fmt.Printf("Also writing summary to %s :\n\n", summaryFileName)
		defer f.Close()
	}

	if filter != "" {
		Printf(f, "# Filtered (%s) Summary:\n\n", filter)
	} else {
		Printf(f, "# Summary:\n\n")
	}
	errorCount, errorString := checkers.FrontSummary(filter)
	Printf(f, errorString)
	count, errorString = checkers.LinkSummary(filter)
	errorCount += count
	//Printf(f, errorString)
	Printf(f, "\n\tFound: %d files\n", len(data.AllFiles))
	Printf(f, "\tFound: %d errors\n", errorCount)
	// return the number of 404's to show that there are things to be fixed
	os.Exit(errorCount)

}

func Printf(f *os.File, format string, a ...interface{}) {
	str := fmt.Sprintf(format, a...)
	fmt.Print(str)
	if f != nil {
		// Don't reall want to know we can't write..
		f.WriteString(str)
	}
}
