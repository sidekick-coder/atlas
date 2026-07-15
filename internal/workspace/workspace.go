package workspace

import (
	"os"
	"path/filepath"
)

var Path string

func Get() (string, error) {
	result := ""

	// command line argument
	if Path != "" {
		result = Path 
	}

	// env variable
	if result == "" {
		envPath := os.Getenv("ATLAS_WORKSPACE_PATH")
		
		if envPath != "" {
			result = envPath
		}
	}

	// cwd
    if result == "" {
		cwdPath, err := os.Getwd()

		if err != nil {
			return "", err
		}

		absPath, err := filepath.Abs(cwdPath)

		result = absPath
    }


    return result, nil
}

func GetRootPath() string {
	rootPath, err := Get()

	if err != nil {
		panic(err)
	}

	return rootPath
}
