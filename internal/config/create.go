package config

import (
	"github.com/sidekick-coder/atlas/internal/workspace"
	"path/filepath"
)

func Create() (*Config, error) {
	path := workspace.GetRootPath()
	entries := map[string]string{}

	entries["workspace.path"] = path
	entries["workspace.atlas_path"] = filepath.Join(path, ".atlas")
	entries["workspace.database_path"] = filepath.Join(path, ".atlas", "database.sqlite")
	
	return New(entries), nil
}
