package config

import (
	"github.com/sidekick-coder/atlas/internal/workspace"
)

func Load() *Config {
	entries := map[string]string{}
	workspacePath, err := workspace.Get()

	if err != nil {
		panic(err)
	}

	entries["workspace.path"] = workspacePath
	
	return New(entries)
}
