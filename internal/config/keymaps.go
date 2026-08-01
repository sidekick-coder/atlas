package config

import (
	"github.com/sidekick-coder/atlas/internal/utils/maputil"
	"github.com/sidekick-coder/atlas/internal/utils/sliceutil"
)

type Keymap struct {
	ID          string         `json:"id"`
	Description string         `json:"description"`
	Keys        []string       `json:"keys"`
	Action      Action         `json:"action"`
	Options     map[string]any `json:"options"`
}

func ConfigKeymapFromMap(m map[string]any) Keymap {
	km := Keymap{}

	if id, ok := m["id"].(string); ok {
		km.ID = id
	}

	if d, ok := m["description"].(string); ok {
		km.Description = d
	}

	if keys, ok := m["keys"].([]any); ok {
		km.Keys = sliceutil.MapString(keys)
	}

	a, err := ParseAction(m["action"])

	if err == nil {
		km.Action = a 
	}

	km.Options = maputil.Except(m, "id", "action", "description", "keys")

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
