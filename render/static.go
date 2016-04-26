package render

import (
	"log"
    "os"
  	 "path/filepath"
)

func CopyStaticFiles(outputDir string, files []string) {
	for _, file := range files {
	log.Println("INFO: <<", file)
	outfile := filepath.FromSlash(filepath.Join(outputDir, file))
        if err := os.MkdirAll(filepath.Dir(outfile), 0755); err != nil {
			log.Println("ERR: ", err)
			continue
        }
}
}
