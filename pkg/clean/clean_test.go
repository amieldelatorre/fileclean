package clean

import (
	"os"
	"path/filepath"
	"reflect"
	"sort"
	"testing"

	"github.com/amieldelatorre/fileclean/pkg/tests"
)

func TestGetDirFileContents(t *testing.T) {
	generatedTestFs, err := tests.GenerateFs(7, 3)
	if err != nil {
		t.Error(err)
	}

	fileContents, err := getDirFileContents(generatedTestFs.TestFsDir)
	if err != nil {
		t.Error(err)
	}

	if len(*fileContents) != generatedTestFs.NumFiles {
		t.Errorf("expected number of files (%d) does not match actual (%d)", generatedTestFs.NumFiles, len(*fileContents))
	}

	os.RemoveAll(generatedTestFs.TestFsDir)
}

func TestGetLabelledFilesAscending(t *testing.T) {
	keep := 3

	generatedTestFs, err := tests.GenerateFs(7, 3)
	if err != nil {
		t.Error(err)
	}

	fileContents, err := getDirFileContents(generatedTestFs.TestFsDir)
	if err != nil {
		t.Error(err)
	}

	expectedKept := generatedTestFs.Files[:keep]
	sortedAscending := GetFileInfoSortedAscending(fileContents)

	actual := getLabelledFiles(sortedAscending, keep, generatedTestFs.TestFsDir)
	actualKept := []string{}
	actualNumToBeKept := 0
	for _, file := range actual {
		if !file.ToBeDeleted {
			actualNumToBeKept++
			actualKept = append(actualKept, filepath.Base(file.Path))
		}
	}

	if actualNumToBeKept != keep {
		t.Errorf("expected number of files to be kept (%d) does not match actual (%d)", keep, actualNumToBeKept)
	}

	if len(*fileContents) - keep != len(actual) - actualNumToBeKept {
		t.Errorf("expected number of files to be deleted (%d) does not match actual (%d)", keep, actualNumToBeKept)
	}

	if !reflect.DeepEqual(actualKept, expectedKept) {
		t.Errorf("expected kept (%v) does not match actual kept (%v)", expectedKept, actualKept)
	}

	os.RemoveAll(generatedTestFs.TestFsDir)
}

func TestGetLabelledFilesDescending(t *testing.T) {
	keep := 3

	generatedTestFs, err := tests.GenerateFs(7, 3)
	if err != nil {
		t.Error(err)
	}

	fileContents, err := getDirFileContents(generatedTestFs.TestFsDir)
	if err != nil {
		t.Error(err)
	}

	sortedDescending := GetFileInfoSortedDescending(fileContents)
	expectedKept := generatedTestFs.Files[generatedTestFs.NumFiles - keep:]

	actual := getLabelledFiles(sortedDescending, keep, generatedTestFs.TestFsDir)
	actualKept := []string{}
	actualNumToBeKept := 0
	for _, file := range actual {
		if !file.ToBeDeleted {
			actualNumToBeKept++
			actualKept = append(actualKept, filepath.Base(file.Path))
		}
	}

	if actualNumToBeKept != keep {
		t.Errorf("expected number of files to be kept (%d) does not match actual (%d)", keep, actualNumToBeKept)
	}

	if len(*fileContents) - keep != len(actual) - actualNumToBeKept {
		t.Errorf("expected number of files to be deleted (%d) does not match actual (%d)", keep, actualNumToBeKept)
	}

	// Sort them here because order doesn't matter during this test, only actual items
	sort.Strings(actualKept)
	sort.Strings(expectedKept)
	if !reflect.DeepEqual(actualKept, expectedKept) {
		t.Errorf("expected kept (%v) does not match actual kept (%v)", expectedKept, actualKept)
	}

	os.RemoveAll(generatedTestFs.TestFsDir)
}

func TestDeleteFilesAscendingDryRun(t *testing.T) {
	keep := 3
	dryRun := true

	generatedTestFs, err := tests.GenerateFs(7, 3)
	if err != nil {
		t.Error(err)
	}

	fileContents, err := getDirFileContents(generatedTestFs.TestFsDir)
	if err != nil {
		t.Error(err)
	}

	sortedAscending := GetFileInfoSortedAscending(fileContents)

	labelledFiles := getLabelledFiles(sortedAscending, keep, generatedTestFs.TestFsDir)
	deleteFiles(labelledFiles, dryRun)


	actualfileContentsAfterDelete, err := getDirFileContents(generatedTestFs.TestFsDir)
	if err != nil {
		t.Error(err)
	}

	actualfileContentsAfterDeleteString := []string{}
	for _, file := range *actualfileContentsAfterDelete{
		actualfileContentsAfterDeleteString = append(actualfileContentsAfterDeleteString, file.Name())
	}

	// Sort them here because order doesn't matter during this test, only actual items
	sort.Strings(actualfileContentsAfterDeleteString)
	sort.Strings(generatedTestFs.Files)
	if !reflect.DeepEqual(actualfileContentsAfterDeleteString, generatedTestFs.Files) {
		t.Errorf("expected kept (%v) does not match actual kept (%v)", actualfileContentsAfterDeleteString, generatedTestFs.Files)
	}

	os.RemoveAll(generatedTestFs.TestFsDir)
}

func TestDeleteFilesAscending(t *testing.T) {
	keep := 3
	dryRun := false

	generatedTestFs, err := tests.GenerateFs(7, 3)
	if err != nil {
		t.Error(err)
	}

	fileContents, err := getDirFileContents(generatedTestFs.TestFsDir)
	if err != nil {
		t.Error(err)
	}

	sortedAscending := GetFileInfoSortedAscending(fileContents)
	expectedKept := generatedTestFs.Files[:keep]

	labelledFiles := getLabelledFiles(sortedAscending, keep, generatedTestFs.TestFsDir)
	deleteFiles(labelledFiles, dryRun)


	actualfileContentsAfterDelete, err := getDirFileContents(generatedTestFs.TestFsDir)
	if err != nil {
		t.Error(err)
	}

	actualfileContentsAfterDeleteString := []string{}
	for _, file := range *actualfileContentsAfterDelete{
		actualfileContentsAfterDeleteString = append(actualfileContentsAfterDeleteString, file.Name())
	}

	// Sort them here because order doesn't matter during this test, only actual items
	sort.Strings(actualfileContentsAfterDeleteString)
	sort.Strings(expectedKept)
	if !reflect.DeepEqual(actualfileContentsAfterDeleteString, expectedKept) {
		t.Errorf("expected kept (%v) does not match actual kept (%v)", actualfileContentsAfterDeleteString, expectedKept)
	}

	os.RemoveAll(generatedTestFs.TestFsDir)
}

func TestDeleteFilesDescending(t *testing.T) {
	keep := 3
	dryRun := false

	generatedTestFs, err := tests.GenerateFs(7, 3)
	if err != nil {
		t.Error(err)
	}

	fileContents, err := getDirFileContents(generatedTestFs.TestFsDir)
	if err != nil {
		t.Error(err)
	}

	sortedDescending := GetFileInfoSortedDescending(fileContents)
	expectedKept := generatedTestFs.Files[generatedTestFs.NumFiles - keep:]

	labelledFiles := getLabelledFiles(sortedDescending, keep, generatedTestFs.TestFsDir)
	deleteFiles(labelledFiles, dryRun)


	actualfileContentsAfterDelete, err := getDirFileContents(generatedTestFs.TestFsDir)
	if err != nil {
		t.Error(err)
	}

	actualfileContentsAfterDeleteString := []string{}
	for _, file := range *actualfileContentsAfterDelete{
		actualfileContentsAfterDeleteString = append(actualfileContentsAfterDeleteString, file.Name())
	}

	// Sort them here because order doesn't matter during this test, only actual items
	sort.Strings(actualfileContentsAfterDeleteString)
	sort.Strings(expectedKept)
	if !reflect.DeepEqual(actualfileContentsAfterDeleteString, expectedKept) {
		t.Errorf("expected kept (%v) does not match actual kept (%v)", actualfileContentsAfterDeleteString, expectedKept)
	}

	os.RemoveAll(generatedTestFs.TestFsDir)
}