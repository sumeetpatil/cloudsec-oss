/*
Copyright © 2022 Sumeet Patil sumeet.patil@sap.com

*/

package search

import (
	"log"
	"os"
	"path/filepath"
)

func Search(clonePath string) []string {
	var files []string

	err := filepath.Walk(clonePath, func(path string, info os.FileInfo, err error) error {

		if err != nil {
			log.Fatalln(err)
			return nil
		}

		if !info.IsDir() && info.Name() == "pom.xml" {
			files = append(files, path)
		}

		//TODO: support more dependencies

		return nil
	})

	if err != nil {
		log.Fatal(err)
	}

	return files
}
