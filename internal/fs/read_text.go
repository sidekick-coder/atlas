package fs

import (
	"os"
	"path/filepath"
)

func ReadText(path string) (string, error) {
	fullPath := filepath.Join(path)

	file, err := os.Open(fullPath)

	if err != nil {
		return "", err
	}

	defer file.Close()

	data, err := os.ReadFile(fullPath)

	if err != nil {
		return "", err
	}

	return string(data), nil
}
