package config

import (
	"fmt"

	"github.com/sidekick-coder/atlas/internal/utils/maputil"
)

type Action struct {
	ID      string         `json:"id"`
	Type    string         `json:"type"`
	Options map[string]any `json:"options"`
}

func ParseAction(payload any) (Action, error) {
	a := Action{}

	if as, ok := payload.(string); ok {
		a.Type = as

		return a, nil
	}

	entry, ok := payload.(map[string]any)

	if !ok {
		return Action{}, fmt.Errorf("invalid action type: %T", payload)
	}

	if id, ok := entry["id"].(string); ok {
		a.ID = id
	}

	if typ, ok := entry["type"].(string); ok {
		a.Type = typ
	}

	a.Options = maputil.Except(entry, "id", "type")

	return a, nil
}

func (c *Config) GetActions() ([]Action, error) {
	entries := c.GetMap("actions")
	actions := []Action{}

	for key, entry := range entries {
		s, err := ParseAction(entry)

		if err != nil {
			return nil, fmt.Errorf("error parsing action entry: %v", err)
		}

		if s.ID == "" {
			s.ID = key
		}

		actions = append(actions, s)
	}

	return actions, nil
}
