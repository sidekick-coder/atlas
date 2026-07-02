package drive 

import (
	"os"
	"github.com/sidekick-coder/atlas/internal/utils"
)

type EntryInfo struct {
	BaseName string
	Path	 string 
	AbsolutePath string
	Type	 string 
	Ext 	 string
}

type Drive struct {
	RootPath string
}

func New(rootPath string) (*Drive, error) {

	isDir, err := utils.IsDirectory(rootPath)

	if err != nil {
		return nil, err
	}

	if !isDir {
		return nil, os.ErrInvalid
	}

	return &Drive{RootPath: rootPath}, nil
}
