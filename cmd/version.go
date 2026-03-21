package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// Version should not be edited directly, will be updated by CI pipeline
const version = "0.0.3"

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Show the version of fileclean",
	Long: `Show the version of fileclean in PATH.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(version)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
