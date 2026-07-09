package config

import "fmt"

type Screen struct {
	ID      string         `json:"id"`
	Type    string         `json:"type"`
	Options map[string]any `json:"options"`
}

func ParseScreen(entry map[string]any) (Screen, error) {
	handler := Screen{}

	if id, ok := entry["id"].(string); ok {
		handler.ID = id
	}

	if typ, ok := entry["type"].(string); ok {
		handler.Type = typ
	}

	options := make(map[string]any)

	for k, v := range entry {
		if k != "id" && k != "type" {
			options[k] = v
		}
	}

	handler.Options = options

	return handler, nil
}

func (c *Config) GetScreens() ([]Screen, error) {
	entries := c.GetArray("screens")
	screens := []Screen{}

	for _, entry := range entries {
		em , ok := entry.(map[string]any)

		if !ok {
			return nil, fmt.Errorf("invalid screen entry: %v", entry)
		}

		s, err := ParseScreen(em)

		if err != nil {
			return nil, fmt.Errorf("error parsing screen entry: %v", err)
		}

		screens = append(screens, s)
	}

	return screens, nil
}
