package config

import (
	"github.com/sidekick-coder/atlas/internal/utils"
)

type Config struct {
	entries map[string]string
}

func New(entries map[string]string) *Config {
	return &Config{entries: entries}
}

func (c *Config) Get(key string) string {
	return c.entries[key]
}

func (c *Config) Set(key, value string) {
	c.entries[key] = value
}

func (c *Config) GetAll() map[string]string {
	return c.entries
}

func (c *Config) GetAllWithPrefix(prefix string) map[string]string {
	result := make(map[string]string)

	for k, v := range c.entries {
		if len(k) >= len(prefix) && k[:len(prefix)] == prefix {
			result[k] = v
		}
	}

	return result
}

func (c *Config) GetAsArray(key string) []any {
	entries := c.GetAllWithPrefix(key)

	input := make(map[string]any)

	for k, v := range entries {
		input[k] = v
	}

	flattend := utils.Unflatten(input)
	result := flattend[key]

	if result == nil {
		return []any{}
	}

	ra, ok := result.([]any)

	if !ok {
		return []any{}
	}

	return ra
}
