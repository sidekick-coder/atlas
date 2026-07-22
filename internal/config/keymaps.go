package config

import (
	"slices"

	"github.com/sidekick-coder/atlas/internal/utils/sliceutil"
)

type Keymap struct {
	ID            string         `json:"id"`
	Description   string         `json:"description"`
	Keys          []string       `json:"keys"`
	Action        string         `json:"action"`
	ActionOptions map[string]any `json:"action_options"`
	Groups        []string       `json:"groups"`
}

func (k Keymap) HasGroup(group ...string) bool {
	for _, g := range group {
		if slices.Contains(k.Groups, g) {
			return true
		}
	}

	return false
}

func ConfigKeymapFromMap(m map[string]any) Keymap {
	km := Keymap{}

	if id, ok := m["id"].(string); ok {
		km.ID = id
	}

	if a, ok := m["action"].(string); ok {
		km.Action = a
	}

	if d, ok := m["description"].(string); ok {
		km.Description = d
	}

	if keys, ok := m["keys"].([]any); ok {
		km.Keys = sliceutil.MapString(keys)
	}

	if groups, ok := m["groups"].([]any); ok {
		km.Groups = sliceutil.MapString(groups)
	}

	if options, ok := m["action_options"].(map[string]any); ok {
		km.ActionOptions = options
	}

	return km
}

func (c *Config) GetKeymaps() []Keymap {
	entries := c.GetMap("keymaps.bindings")
	keymaps := make([]Keymap, 0)

	for key, v := range entries {
		vm, ok := v.(map[string]any)

		if !ok {
			continue
		}

		k := ConfigKeymapFromMap(vm)

		if k.ID == "" {
			k.ID = key
		}

		keymaps = append(keymaps, k)
	}

	return keymaps
}

func (c *Config) GetKeymapsByGroup(group string) []Keymap {
	all := c.GetKeymaps()

	filtered := make([]Keymap, 0)

	for _, keymap := range all {
		if slices.Contains(keymap.Groups, group) {
			filtered = append(filtered, keymap)
		}
	}

	return filtered
}
