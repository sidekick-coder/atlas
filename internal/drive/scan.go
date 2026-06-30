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

		RelativePath, err := filepath.Rel(root, path)

		entries = append(entries, Entry{
			Path:  RelativePath,
			Name:  d.Name(),
			IsDir: d.IsDir(),
		})

		return nil
	})

	return entries, err
}
