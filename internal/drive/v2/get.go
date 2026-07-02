package drive 

import (
	"os"
	"io/fs"
	"path/filepath"
)

func (d *Drive) Get(path string) (*EntryInfo, error) {
	absolutePath := filepath.Join(d.RootPath, path)

	info, err := fs.Stat(fs.FS(os.DirFS(d.RootPath)), path)

	if err != nil {
		return nil, err
	}

	entryType := "file"

	if info.IsDir() {
		entryType = "directory"
	}

	entryInfo := EntryInfo{
		AbsolutePath: absolutePath,
		Path:         path,
		Type:         entryType,
		Ext:          filepath.Ext(info.Name()),
		BaseName:     info.Name(),
	}

	return &entryInfo, nil
}
