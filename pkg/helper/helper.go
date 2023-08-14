package helper

import (
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
)

//Check if directory exist or is not empty
func IsEmptyDir(path string) (bool, error) {
	f, err := os.Open(path)
	if err != nil {
		return false, err
	}
	defer f.Close()

	_, err = f.Readdirnames(1)
	if err == io.EOF {
		return true, nil
	}
	return false, err
}

//Check if path is valid
func IsFilePathValid(path string) bool {
	if _, err := os.Stat(path); err == nil {
		return true
	}

	var d []byte
	if err := ioutil.WriteFile(path, d, 0644); err == nil {
		os.Remove(path)
		return true
	}

	return false
}

func GetAllFilesInDirectory(path string) ([]string, error) {
	var files []string
	err := filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			files = append(files, path)
		}
		return nil
	})

	return files, err
}
