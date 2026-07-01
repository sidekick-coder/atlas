package drive

import (
	"io/fs"
	"path/filepath"
)

func ScanStream(root string, fn func(Entry) error, options ...ScanOptions) error {
	var ignores []string

	if len(options) > 0 {
		ignores = options[0].Ignores
	}

	patterns := CreateIgnorePatterns(ignores...)

	return filepath.WalkDir(root, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		RelativePath, err := filepath.Rel(root, path)

		if (RelativePath == "." || RelativePath == "..") && d.IsDir() {
			return nil
		}

		if ShouldIgnore(RelativePath, patterns) {
			if d.IsDir() {
				return filepath.SkipDir
			}

			return nil
		}

		return fn(Entry{
			Path:  RelativePath,
			Name:  d.Name(),
			IsDir: d.IsDir(),
		})
	})
}
