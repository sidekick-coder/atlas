package config

import (
	"path/filepath"
	"github.com/sidekick-coder/atlas/internal/fs"
	"github.com/sidekick-coder/atlas/internal/utils"
	"github.com/sidekick-coder/atlas/internal/workspace"
)

func Create() (*Config, error) {
	path := workspace.GetRootPath()
	entries := map[string]string{}

	if (fs.Exists(filepath.Join(path, ".atlas", "config.yml"))) {
		yaml, err := fs.ReadYaml(filepath.Join(path, ".atlas", "config.yml"))
		
		if err != nil {
			return nil, err
		}

		flattened := utils.FlattenMap(yaml, "")

		flatStrings := utils.StringifyMap(flattened)

		for k, v := range flatStrings {
			entries[k] = v
		}
	}

	entries["workspace.path"] = path
	entries["workspace.atlas_path"] = filepath.Join(path, ".atlas")
	entries["workspace.database_path"] = filepath.Join(path, ".atlas", "database.sqlite")
	
	return New(entries), nil
}
