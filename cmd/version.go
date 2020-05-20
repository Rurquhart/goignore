package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var version string

func init() {
	version = "v1.0.0"
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use: "version",
	Short: "Prints the installs version of goignore",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("goignore %s", version)
	},
}