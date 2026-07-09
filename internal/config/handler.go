package config

import "fmt"

type Handler struct {
	ID      string         `json:"id"`
	Type    string         `json:"type"`
	Patterns []string         `json:"patterns"`
	Options map[string]any `json:"options"`
}

func ParseHandler(entry map[string]any) (Handler, error) {
	handler := Handler{}

	if id, ok := entry["id"].(string); ok {
		handler.ID = id
	}

	if typ, ok := entry["type"].(string); ok {
		handler.Type = typ
	}

	if patterns, ok := entry["patterns"].([]any); ok {
		for _, p := range patterns {
			if ps, ok := p.(string); ok {
				handler.Patterns = append(handler.Patterns, ps)
			} else {
				return handler, fmt.Errorf("invalid pattern: %v", p)
			}
		}
	}

	options := make(map[string]any)

	for k, v := range entry {
		if k != "id" && k != "type" && k != "patterns" {
			options[k] = v
		}
	}

	handler.Options = options

	return handler, nil
}

func (c *Config) GetConfigHandlers() ([]Handler, error) {
	entries := c.GetArray("handlers")
	handlers := []Handler{}

	for _, entry := range entries {
		em , ok := entry.(map[string]any)

		if !ok {
			return nil, fmt.Errorf("invalid handler entry: %v", entry)
		}

		h, err := ParseHandler(em)

		if err != nil {
			return nil, fmt.Errorf("error parsing handler entry: %v", err)
		}

		handlers = append(handlers, h)

	}

	return handlers, nil
}
