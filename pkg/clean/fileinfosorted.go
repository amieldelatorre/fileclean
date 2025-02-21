package clean

import (
	"os"
	"sort"
)

type FileInfoSorted interface {
	Values() []os.FileInfo
	Len() int
	Less(i, j int) bool
	Swap(i, j int)
}

type FileInfoSortedAscending []os.FileInfo

func (p FileInfoSortedAscending) Values() []os.FileInfo {
	return p
}

func (p FileInfoSortedAscending) Len() int {
	return len(p)
}

func (p FileInfoSortedAscending) Less(i, j int) bool {
	return p[i].ModTime().Before(p[j].ModTime())
}

func (p FileInfoSortedAscending) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

type FileInfoSortedDescending []os.FileInfo

func (p FileInfoSortedDescending) Values() []os.FileInfo {
	return p
}

func (p FileInfoSortedDescending) Len() int {
	return len(p)
}

func (p FileInfoSortedDescending) Less(i, j int) bool {
	return p[i].ModTime().After(p[j].ModTime())
}

func (p FileInfoSortedDescending) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func GetFileInfoSortedAscending(files *[]os.FileInfo) FileInfoSortedAscending {
	sortedFiles := make(FileInfoSortedAscending, 0, len(*files))
	for _, file := range *files {
		sortedFiles = append(sortedFiles, file)
	}
	sort.Sort(sortedFiles)
	return sortedFiles
}

func GetFileInfoSortedDescending(files *[]os.FileInfo) FileInfoSortedDescending {
	sortedFiles := make(FileInfoSortedDescending, 0, len(*files))
	for _, file := range *files {
		sortedFiles = append(sortedFiles, file)
	}
	sort.Sort(sortedFiles)
	return sortedFiles
}
