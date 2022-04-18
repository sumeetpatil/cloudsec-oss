/*
Copyright © 2022 Sumeet Patil sumeet.patil@sap.com

*/
package cmd

import (
	"bytes"
	"cloudsec-oss/internal/parse"
	"cloudsec-oss/internal/search"
	"cloudsec-oss/internal/sourcecontrol"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/spf13/cobra"
)

var scanCmd = &cobra.Command{
	Use:   "scan",
	Short: "Scan a git repository for open source vulnerabilites",
	Long: `Scan a git repository for open source vulnerabilites
	For Example: scan --git https://github.com/sumeetpatil/HibenateSpringExample`,
	Run: func(cmd *cobra.Command, args []string) {
		git, _ := cmd.Flags().GetString("git")
		tmpDir := sourcecontrol.Clone(git)
		defer os.RemoveAll(tmpDir)

		files := search.Search(tmpDir)

		for _, file := range files {
			pomDependencies := parse.ParsePomDependencies(file)
			for _, dependency := range pomDependencies {
				data := "{\"version\": \"" + dependency.Version + "\", \"package\": {\"name\": \"" + dependency.LibraryName + "\", \"ecosystem\": \"" + dependency.LibraryType + "\"}}"
				resp, err := http.Post("https://api.osv.dev/v1/query", "application/json", bytes.NewBuffer([]byte(data)))
				if err != nil {
					log.Fatalln(err)
					continue
				}

				defer resp.Body.Close()

				body, err := ioutil.ReadAll(resp.Body)

				if err != nil {
					log.Fatalln(err)
					continue
				}

				var result parse.Response
				if err := json.Unmarshal([]byte(string(body)), &result); err != nil {
					log.Fatalln("Can not unmarshal JSON")
					log.Fatalln(err)
					continue
				}

				if result.Vulns != nil {
					preetyJson, _ := json.MarshalIndent(result.Vulns, "", "\t")
					log.Println(string(preetyJson))
				}
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(scanCmd)
	scanCmd.Flags().StringP("git", "g", "", "Git URL")
}
