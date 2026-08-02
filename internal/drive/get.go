package drive 

import (
	"os"
	"io/fs"
	"github.com/sidekick-coder/atlas/internal/models"
)

func (d *Drive) Exists(path string) (bool, error) {
	_, err := fs.Stat(fs.FS(os.DirFS(d.path)), path)

	if err != nil {
		if os.IsNotExist(err) {
			return false, nil
		}

		return false, err
	}

	return true, nil
}

func (d *Drive) Get(path string) (*models.EntryInfo, error) {
	info, err := fs.Stat(fs.FS(os.DirFS(d.path)), path)

	if err != nil {
		return nil, err
	}

	entryInfo := models.NewEntryInfoFromFileInfo(d.path, path, info)

	return &entryInfo, nil
}
