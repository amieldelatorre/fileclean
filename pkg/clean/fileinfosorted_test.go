package clean

import (
	"os"
	"testing"

	"github.com/amieldelatorre/fileclean/pkg/tests"
)

func TestGetFileInfoSortedAscending(t *testing.T) {
	generatedTestFs, err := tests.GenerateFs(7, 3)
	if err != nil {
		t.Error(err)
	}
	
	fileContents, err := getDirFileContents(generatedTestFs.TestFsDir)
	if err != nil {
		t.Error(err)
	}

	actual := GetFileInfoSortedAscending(fileContents)
	if len(actual.Values()) != generatedTestFs.NumFiles {
		t.Errorf("expected number of sorted files (%d) does not match actual (%d)", generatedTestFs.NumFiles, len(actual.Values()))
	}

	for i := 0; i < generatedTestFs.NumFiles; i++ {
		actualFileName := actual.Values()[i].Name()
		expectedFileName := generatedTestFs.Files[i]
		if actualFileName != expectedFileName {
			t.Errorf("expected name of sorted files (%s) does not match actual (%s)", expectedFileName, actualFileName)
		}
	}

	os.RemoveAll(generatedTestFs.TestFsDir)
}

func TestGetFileInfoSortedDescending(t *testing.T) {
	generatedTestFs, err := tests.GenerateFs(7, 3)
	if err != nil {
		t.Error(err)
	}
	
	fileContents, err := getDirFileContents(generatedTestFs.TestFsDir)
	if err != nil {
		t.Error(err)
	}

	actual := GetFileInfoSortedDescending(fileContents)
	if len(actual.Values()) != generatedTestFs.NumFiles {
		t.Errorf("expected number of sorted files (%d) does not match actual (%d)", generatedTestFs.NumFiles, len(actual.Values()))
	}

	for i := 0; i < generatedTestFs.NumFiles; i++ {
		actualFileName := actual.Values()[i].Name()
		expectedFileName := generatedTestFs.Files[generatedTestFs.NumFiles - i - 1]
		if actualFileName != expectedFileName {
			t.Errorf("expected name of sorted files (%s) does not match actual (%s)", expectedFileName, actualFileName)
		}
	}

	os.RemoveAll(generatedTestFs.TestFsDir)
}