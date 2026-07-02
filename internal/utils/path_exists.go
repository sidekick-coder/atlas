package utils 

import (
	"os"
)

func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)

	if err != nil {
		return false, err
	}

	return true, nil
}

func IsDirectory(path string) (bool, error) {
	info, err := os.Stat(path)

	if err != nil {
		return false, err
	}

	return info.IsDir(), nil
}
