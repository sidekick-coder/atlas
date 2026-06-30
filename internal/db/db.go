package db

import (
	"os"
	"path/filepath"
)

func Path(workspace string) (string, error) {
	atlasDir := filepath.Join(workspace, ".atlas")

	if err := os.MkdirAll(atlasDir, 0755); err != nil {
		return "", err
	}

	return filepath.Join(atlasDir, "database.sqlite"), nil
}
