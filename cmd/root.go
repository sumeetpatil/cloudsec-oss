/*
Copyright © 2022 Sumeet Patil sumeet.patil@sap.com

*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "cloudsec-oss",
	Short: "Cloud Security for opensource software",
	Long:  `Cloud Security for opensource software`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
}
