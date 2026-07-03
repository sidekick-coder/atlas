package drive 

import (
	"os"
	"io/fs"
	"github.com/sidekick-coder/atlas/internal/models"
)

func (d *Drive) Get(path string) (*models.EntryInfo, error) {
	info, err := fs.Stat(fs.FS(os.DirFS(d.RootPath)), path)

	if err != nil {
		return nil, err
	}

	entryInfo := models.NewEntryInfoFromFileInfo(d.RootPath, path, info)

	return &entryInfo, nil
}
