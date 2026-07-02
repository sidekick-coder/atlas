package workspace 

import (
	"os"
	"path/filepath"
)

var Path string

func Get() (string, error) {
    if Path != "" {
        return filepath.Abs(Path)
    }

    return os.Getwd()
}

func GetRootPath() string {
	rootPath, err := Get()

	if err != nil {
		panic(err)
	}

	return rootPath
}
