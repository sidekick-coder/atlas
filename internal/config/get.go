package config

import (
	"fmt"
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

func (c *Config) GetMap(key string) map[string]any {
	entries := c.GetByPrefix(key)

	input := make(map[string]any)

	for k, v := range entries {
		input[k] = v
	}

	flattend := utils.Unflatten(input)
	result := utils.Get(flattend, key)

	if result == nil {
		return map[string]any{}
	}

	// convert array of maps to map 
	if arr, ok := result.([]any); ok {
		rm := map[string]any{}

		for i, v := range arr {
			if m, ok := v.(map[string]any); ok {
				rm[fmt.Sprintf("%d", i)] = m
			}
		}

		return rm
	}

	rm, ok := result.(map[string]any)

	if !ok {
		return map[string]any{}
	}

	return rm
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
