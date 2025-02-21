package cmd

import (
	"github.com/amieldelatorre/fileclean/pkg/clean"
	"github.com/amieldelatorre/fileclean/pkg/sortorder"
	"github.com/spf13/cobra"
)

const (
	pathFlagName      = "path"
	keepFlagName      = "keep"
	sortOrderFlagName = "sort-order"
	dryRunFlagName    = "dry-run"
)

var (
	cleanPath = ""
	keep      = -1
	sortOrder = sortorder.SortOrderDescending
	dryRun    = false
)

var cleanCmd = &cobra.Command{
	Use:   "clean",
	Short: "Runs the file cleaning.",
	Long: `Runs the file cleaning.

	Removes files and keeps x amount based on sorting and desired amount.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		return clean.Execute(cleanPath, keep, sortOrder.String(), dryRun)
	},
}

func init() {
	cleanCmd.Flags().StringVar(&cleanPath, pathFlagName, "", "Directory that is being cleaned")
	cleanCmd.MarkFlagRequired("path")

	cleanCmd.Flags().IntVar(&keep, keepFlagName, 5, "Keeps the first x number of files based on sort order")
	cleanCmd.MarkFlagRequired("keep")

	cleanCmd.Flags().Var(&sortOrder, sortOrderFlagName, "Sort order of the files, ascending or descending order")
	cleanCmd.MarkFlagRequired(sortOrderFlagName)
	cleanCmd.RegisterFlagCompletionFunc(sortOrderFlagName, sortorder.SortOrderCompletion)

	cleanCmd.Flags().BoolVar(&dryRun, dryRunFlagName, false, "Show the files to be deleted without actually deleting them")

	rootCmd.AddCommand(cleanCmd)
}
