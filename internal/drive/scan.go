package drive

import (
	"io/fs"
	"path/filepath"
	"strings"

	"github.com/sidekick-coder/atlas/internal/models"
)

type ScanOptions struct {
	Ignores []string
}

func (d *Drive) shouldDescend(dir string) bool {
    dir = filepath.Clean(dir)
	includes := d.config.GetArrayString("scan.include")

    for _, inc := range includes {
        inc = filepath.Clean(inc)

        if strings.HasPrefix(inc, dir+string(filepath.Separator)) {
            return true
        }
    }

    return false
}

func (d *Drive) ScanStream(fn func(models.EntryInfo) error, options ...ScanOptions) error {
	ignores := []string{}
	root := d.path

	if len(options) > 0 {
		ignores = options[0].Ignores
	}

	return filepath.WalkDir(root, func(path string, dirEntry fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		RelativePath, err := filepath.Rel(root, path)

		if (RelativePath == "." || RelativePath == "..") && dirEntry.IsDir() {
			return nil
		}

		ignore, err := d.Ignore(RelativePath, ignores...) 

		if err != nil {
			return err
		}


		if ignore {
			if !dirEntry.IsDir() {
				return nil
			}

			if !d.shouldDescend(RelativePath) {
				return filepath.SkipDir
			}

			return nil
		}

		entryInfo := models.NewEntryInfoFromDirEntry(root, RelativePath, dirEntry)

		return fn(entryInfo)
	})
}

func (d *Drive) Scan(options ...ScanOptions) ([]models.EntryInfo, error) {
	var entries []models.EntryInfo

	cb := func(entry models.EntryInfo) error {
		entries = append(entries, entry)
		return nil
	}

	err := d.ScanStream(cb, options...)


	return entries, err
}

