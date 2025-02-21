package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "fileclean",
	Short: "Cleans a directory based on flags.",
	Long: `Cleans a directory based on flags.
	
	Removes files and keeps x amount based on sorting and desired amount.`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.SilenceUsage = true
}
