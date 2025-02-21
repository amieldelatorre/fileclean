package tests

import (
	"fmt"
	"os"
	"path/filepath"
	"time"
)

const (
	testDataDirPrefix = "_testdata"
)

type GeneratedFs struct {
	Files     []string
	NumFiles  int
	Dirs      []string
	NumDirs   int
	TestFsDir string
}

func GenerateFs(numFiles int, numDirs int) (*GeneratedFs, error) {
	generatedFs := GeneratedFs{}
	testDataDir, err := os.MkdirTemp(".", testDataDirPrefix)
	if err != nil {
		return nil, err
	}

	files := []string{}
	for i := 0; i < numFiles; i++ {
		file, err := os.CreateTemp(testDataDir, fmt.Sprintf("file%d_", i))
		if err != nil {
			return nil, err
		}
		files = append(files, filepath.Base(file.Name()))

		err = file.Close()
		if err != nil {
			return nil, err
		}
		
		time.Sleep(time.Millisecond * 100)
	}

	dirs := []string{}
	for i := 0; i < numDirs; i++ {
		dir, err := os.MkdirTemp(testDataDir, fmt.Sprintf("dir%d_", i))
		if err != nil {
			return nil, err
		}
		dirs = append(dirs, dir)
	}

	generatedFs.Files = files
	generatedFs.NumFiles = len(files)
	generatedFs.Dirs = dirs
	generatedFs.NumDirs = len(dirs)
	generatedFs.TestFsDir = testDataDir
	return &generatedFs, nil
}
