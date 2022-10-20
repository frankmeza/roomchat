package utils

import (
	"path/filepath"
	"runtime"
)

var (
	_, rootDir, _, _ = runtime.Caller(0)

	// Root folder of this project
	rootDirectory = filepath.Join(filepath.Dir(rootDir), "../..")
)

func GetRootPath(path string) string {
	return rootDirectory + "/" + path
}
