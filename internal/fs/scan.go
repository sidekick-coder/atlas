package fs

import (
	"io/fs"
	"path/filepath"
)

func Scan(path string) ([]string, error) {
	entries := []string{}

	filepath.WalkDir(path, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if d.IsDir() {
			return nil
		}

		if filepath.Ext(path) != ".yaml" && filepath.Ext(path) != ".yml" {
			return nil
		}

		entries = append(entries, path)

		return nil
	})

	return entries, nil
}
