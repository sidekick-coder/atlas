package config

import (
	"strings"

	"github.com/sidekick-coder/atlas/internal/utils"
)

func (c *Config) GetAll() map[string]string {
	return c.entries
}

func (c *Config) Get(key string) (string, bool) {
	v, ok := c.entries[key]

	if !ok {
		return "", false
	}

	return v, true
}


func (c *Config) GetByPrefix(prefix string) map[string]string {
	result := make(map[string]string)

	for k, v := range c.entries {
		if strings.HasPrefix(k, prefix) {
			result[k] = v
		}
	}


	return result
}

func (c *Config) GetArray(key string) []any {
	entries := c.GetByPrefix(key)

	input := make(map[string]any)

	for k, v := range entries {
		input[k] = v
	}

	flattend := utils.Unflatten(input)
	result := utils.Get(flattend, key)

	if result == nil {
		return []any{}
	}

	ra, ok := result.([]any)

	if !ok {
		return []any{}
	}

	return ra
}

func (c *Config) GetArrayString(key string) []string {
	entries := c.GetArray(key)


	result := []string{}

	for _, v := range entries {
		if str, ok := v.(string); ok {
			result = append(result, str)
		}
	}

	return result
}
