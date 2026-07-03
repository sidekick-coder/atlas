package drive

import (
	"io/fs"
	"path/filepath"
	"github.com/sidekick-coder/atlas/internal/models"
)

type ScanOptions struct {
	Ignores []string
}

func (d *Drive) ScanStream(fn func(models.EntryInfo) error, options ...ScanOptions) error {
	ignores := []string{}
	root := d.RootPath

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

		entryInfo := models.NewEntryInfoFromDirEntry(root, RelativePath, d)

		return fn(entryInfo)
	})
}
