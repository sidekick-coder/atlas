package drive

import (
	"io/fs"
	"path/filepath"
)

func Scan(root string) ([]Entry, error) {
	var entries []Entry

	err := filepath.WalkDir(root, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		entries = append(entries, Entry{
			Path:  path,
			Name:  d.Name(),
			IsDir: d.IsDir(),
		})

		return nil
	})

	return entries, err
}
