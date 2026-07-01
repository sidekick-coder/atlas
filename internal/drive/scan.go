package drive

import (
	"io/fs"
	"path/filepath"
)

type ScanOptions struct {
	Ignores []string
}

func Scan(root string, options ...ScanOptions) ([]Entry, error) {
	var ignores []string

	if len(options) > 0 {
		ignores = options[0].Ignores
	}

	patterns := CreateIgnorePatterns(ignores...)

	var entries []Entry

	err := filepath.WalkDir(root, func(path string, d fs.DirEntry, err error) error {
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


		entries = append(entries, Entry{
			Path:  RelativePath,
			Name:  d.Name(),
			IsDir: d.IsDir(),
		})

		return nil
	})

	return entries, err
}
