package config

import "fmt"

type Action struct {
	ID      string         `json:"id"`
	Type    string         `json:"type"`
	Options map[string]any `json:"options"`
}

func ParseAction(entry map[string]any) (Action, error) {
	a := Action{}

	if id, ok := entry["id"].(string); ok {
		a.ID = id
	}

	if typ, ok := entry["type"].(string); ok {
		a.Type = typ
	}

	options := make(map[string]any)

	for k, v := range entry {
		if k != "id" && k != "type" {
			options[k] = v
		}
	}

	a.Options = options

	return a, nil
}

func (c *Config) GetActions() ([]Action, error) {
	entries := c.GetMap("actions")
	actions := []Action{}

	for key, entry := range entries {
		em, ok := entry.(map[string]any)

		if !ok {
			return nil, fmt.Errorf("invalid action entry: %v", entry)
		}

		s, err := ParseAction(em)

		if s.ID == "" {
			s.ID = key
		}

		if err != nil {
			return nil, fmt.Errorf("error parsing screen entry: %v", err)
		}

		actions = append(actions, s)
	}

	return actions, nil
}
