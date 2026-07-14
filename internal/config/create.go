package config

import (
	"maps"
	"path/filepath"
	"strings"

	"github.com/sidekick-coder/atlas/internal/fs"
	"github.com/sidekick-coder/atlas/internal/utils"
	"github.com/sidekick-coder/atlas/internal/workspace"
)

func Create() (*Config, error) {
	path := workspace.GetRootPath()

	c := &Config{
		entries: map[string]string{},
	}

	c.entries["workspace.database_path"] = filepath.Join(path, ".atlas", "database.sqlite")

	err := c.LoadFiles([]string{
		filepath.Join(path, ".atlas", "config.yaml"),
		filepath.Join(path, ".atlas", "config.yml"),
	})

	if err != nil {
		return nil, err
	}

	err = c.LoadFolder(filepath.Join(path, ".atlas", "config"))

	if err != nil {
		return nil, err
	}

	c.entries["workspace.path"] = path
	c.entries["workspace.atlas_path"] = filepath.Join(path, ".atlas")

	return c, nil
}

func (c *Config) LoadFile(filename string, prefixes ...string) error {
	if !fs.Exists(filename) {
		return nil
	}

	yaml, err := fs.ReadYaml(filename)

	if err != nil {
		return err
	}

	prefix := strings.Join(prefixes, ".")

	f := utils.FlattenMap(yaml, prefix)
	fs := utils.StringifyMap(f)

	maps.Copy(c.entries, fs)

	return nil
}

func (c *Config) LoadFiles(filenames []string, prefixes ...string) error {
	for _, filename := range filenames {
		err := c.LoadFile(filename, prefixes...)

		if err != nil {
			return err
		}
	}

	return nil
}

func (c *Config) LoadFolder(folder string) error {
	if !fs.Exists(folder) {
		return nil
	}

	entries, err := fs.Scan(folder)

	if err != nil {
		return err 
	}

	for _, entry := range entries {
		relPath, err := filepath.Rel(folder, entry)

		if err != nil {
			return err
		}

		prefix := strings.TrimSuffix(relPath, filepath.Ext(relPath))

		prefix = strings.ReplaceAll(prefix, string(filepath.Separator), ".")

		err = c.LoadFile(entry, prefix)
	}

	return nil
}
