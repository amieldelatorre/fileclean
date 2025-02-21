package clean

import (
	"errors"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"time"

	"github.com/amieldelatorre/fileclean/pkg"
	"github.com/amieldelatorre/fileclean/pkg/sortorder"
)

type LabelledFile struct {
	ToBeDeleted bool
	Path        string
	ModTime     time.Time
}

func Execute(cleanPath string, keep int, sortOrder string, dryRun bool) error {
	pathInfo, err := os.Stat(cleanPath)
	if err == nil && !pathInfo.IsDir() {
		pkg.LogErrorExit("%s is not a directory", cleanPath)
	} else if err != nil && errors.Is(err, fs.ErrNotExist) {
		fmt.Println("path does not exist")
		os.Exit(1)
	} else if err != nil {
		panic(err)
	}

	files, err := getDirFileContents(cleanPath)
	if err != nil {
		return err
	}

	var sortedFiles FileInfoSorted
	switch sortOrder {
	case string(sortorder.SortOrderAscending):
		sortedFiles = GetFileInfoSortedAscending(files)
	case string(sortorder.SortOrderDescending):
		sortedFiles = GetFileInfoSortedDescending(files)
	default:
		return errors.New("unknown sort-order")
	}

	markedFiles := getLabelledFiles(sortedFiles, keep, cleanPath)
	err = deleteFiles(markedFiles, dryRun)
	return err
}

func getDirFileContents(cleanPath string) (*[]os.FileInfo, error) {
	files := []os.FileInfo{}

	dirContents, err := os.ReadDir(cleanPath)
	if err != nil {
		return &files, err
	}

	for _, item := range dirContents {
		if !item.IsDir() {
			fileInfo, err := item.Info()
			if err != nil {
				return nil, err
			}

			files = append(files, fileInfo)
		}
	}

	return &files, nil
}

func getLabelledFiles(files FileInfoSorted, keep int, cleanPath string) []LabelledFile {
	labelledFiles := []LabelledFile{}
	count := 1
	for _, file := range files.Values() {
		toBeDeleted := false
		if count > keep {
			toBeDeleted = true
		}

		labelledFile := LabelledFile{
			Path:        filepath.Join(cleanPath, file.Name()),
			ModTime:     file.ModTime(),
			ToBeDeleted: toBeDeleted,
		}

		labelledFiles = append(labelledFiles, labelledFile)
		count += 1
	}

	return labelledFiles
}

func deleteFiles(files []LabelledFile, dryRun bool) error {
	for _, file := range files {
		if !file.ToBeDeleted {
			pkg.LogInformation("%s   "+file.Path, "SKIPPING")
			continue
		}

		if dryRun {
			pkg.LogWarn("%s     "+file.Path, "DELETE")
			continue
		}

		if !dryRun {
			err := os.Remove(file.Path)
			if err != nil {
				return err
			}

			pkg.LogWarn("%s    "+file.Path, "DELETED")
		}

	}

	return nil
}
