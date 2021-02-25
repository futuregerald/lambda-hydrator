package main

import (
	"log"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/karrick/godirwalk"
)

func main() {
	dirName, err := filepath.Abs("../functions_src")
	if err != nil {
		log.Print(err)
	}
	err = godirwalk.Walk(dirName, &godirwalk.Options{
		Callback: func(osPathname string, de *godirwalk.Dirent) error {
			var functionsDir []string
			allDir := strings.Replace(osPathname, dirName, "", 1)
			if runtime.GOOS == "windows" {
				functionsDir = strings.Split(allDir, `\`)
			} else {
				functionsDir = strings.Split(allDir, `/`)
			}
			if len(functionsDir) > 2 {
				return godirwalk.SkipThis
			}
			if len(functionsDir) == 2 {
				log.Print(functionsDir)
			}

			return nil
		},
		Unsorted: true,
	})
	if err != nil {
		log.Print(err)
	}
}
