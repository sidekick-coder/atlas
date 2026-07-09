package drive

import (
	"os"

	"github.com/sidekick-coder/atlas/internal/config"
	"github.com/sidekick-coder/atlas/internal/utils"
)

type Drive struct {
	path string
	config   *config.Config
}

func New(path string) (*Drive, error) {
	isDir, err := utils.IsDirectory(path)

	if err != nil {
		return nil, err
	}

	if !isDir {
		return nil, os.ErrInvalid
	}

	return &Drive{path: path}, nil
}

func CreateFromConfig(config *config.Config) (*Drive, error) {
	path, ok := config.Get("workspace.path")

	if !ok {
		return nil, os.ErrInvalid
	}

	if path == "" {
		return nil, os.ErrInvalid
	}

	d, err := New(path)

	if err != nil {
		return nil, err
	}

	d.SetConfig(config)

	return d, nil
}

func (d *Drive) SetConfig(config *config.Config) *Drive {
	d.config = config
	return d
}

