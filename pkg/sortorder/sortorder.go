package sortorder

import (
	"errors"

	"github.com/spf13/cobra"
)

type SortOrder string

func (e *SortOrder) String() string {
	return string(*e)
}

func (e *SortOrder) Set(v string) error {
	switch v {
	case string(SortOrderAscending), string(SortOrderDescending):
		*e = SortOrder(v)
		return nil
	default:
		return errors.New(`must be one of "ascending", or "descending"`)
	}
}

func (e *SortOrder) Type() string {
	return "SortOrder"
}

func SortOrderCompletion(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
	return []string{
		"latest",
		"oldest",
	}, cobra.ShellCompDirectiveDefault
}

const (
	SortOrderAscending  SortOrder = "ascending"
	SortOrderDescending SortOrder = "descending"
)
