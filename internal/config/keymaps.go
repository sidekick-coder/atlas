package config

import (
	"slices"
)

type ConfigKeymap struct {
	Description string   `json:"description"`
	Keys        []string `json:"keys"`
	Action      string   `json:"action"`
	Groups      []string `json:"groups"`
}

func toStringSlice(v any) []string {
	arr, ok := v.([]any)
	if !ok {
		return nil
	}

	out := make([]string, 0, len(arr))
	for _, x := range arr {
		if s, ok := x.(string); ok {
			out = append(out, s)
		}
	}

	return out
}

func ConfigKeymapFromMap(m map[string]any) ConfigKeymap {
	keys := []string{}
	action := ""
	groups := []string{}
	description := ""

	keys = toStringSlice(m["keys"])
	groups = toStringSlice(m["groups"])

	a, ok := m["action"].(string)

	if ok {
		action = a
	}

	d, ok := m["description"].(string)

	if ok {
		description = d
	}

	return ConfigKeymap{
		Description: description,
		Keys:        keys,
		Action:      action,
		Groups:      groups,
	}
}

func (c *Config) GetKeymaps() []ConfigKeymap {
	entries := c.GetAsArray("keymaps")
	keymaps := make([]ConfigKeymap, 0)

	for _, v := range entries {
		vm, ok := v.(map[string]any)

		if !ok {
			continue
		}

		k := ConfigKeymapFromMap(vm)

		keymaps = append(keymaps, k)
	}

	return keymaps
}

func (c *Config) GetKeymapsByGroup(group string) []ConfigKeymap {
	all := c.GetKeymaps()

	filtered := make([]ConfigKeymap, 0)

	for _, keymap := range all {
		if slices.Contains(keymap.Groups, group) {
			filtered = append(filtered, keymap)
		}
	}

	return filtered
}
