/*
Copyright © 2022 Sumeet Patil sumeet.patil@sap.com

*/

package sourcecontrol

import (
	"io/ioutil"
	"log"

	"github.com/go-git/go-git/v5"
)

func Clone(url string) string {
	tmpDir, err := ioutil.TempDir("tmp", "oss")
	if err != nil {
		log.Fatal(err)
	}

	if err != nil {
		log.Fatal(err)
	}

	r, err := git.PlainClone(tmpDir, false, &git.CloneOptions{
		URL: url,
	})

	if err != nil {
		log.Fatal(err)
	}

	log.Println("Cloned ", r)

	return tmpDir
}
