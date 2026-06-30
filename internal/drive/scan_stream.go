package drive

import (
	"io/fs"
	"path/filepath"
)

func ScanStream(root string, fn func(Entry) error) error {
	return filepath.WalkDir(root, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		RelativePath, err := filepath.Rel(root, path)

		return fn(Entry{
			Path:  RelativePath,
			Name:  d.Name(),
			IsDir: d.IsDir(),
		})
	})
}
