package utils

import (
	"os"

	radicalLogger "github.com/W3-Engineers-Ltd/Radical/logger"
)

// Mkdir ...
func Mkdir(dir string) bool {
	if dir == "" {
		radicalLogger.Log.Fatalf("The directory is empty")
		return false
	}
	err := os.MkdirAll(dir, 0755)
	if err != nil {
		radicalLogger.Log.Fatalf("Could not create the directory: %s", err)
		return false
	}

	radicalLogger.Log.Infof("Create %s Success!", dir)
	return true
}

// IsDir ...
func IsDir(dir string) bool {
	f, e := os.Stat(dir)
	if e != nil {
		return false
	}
	return f.IsDir()
}

// IsExist returns whether a file or directory exists.
func IsExist(path string) bool {
	_, err := os.Stat(path)
	return err == nil || os.IsExist(err)
}
