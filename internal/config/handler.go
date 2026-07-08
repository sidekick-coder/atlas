package config

// import "fmt"

// "slices"

type ConfigHandler struct {
	ID      string         `json:"id"`
	Type    string         `json:"type"`
	Pattern string         `json:"pattern"`
	Options map[string]any `json:"options"`
}

func (c *Config) GetConfigHandlers() []ConfigHandler {
	// entries := c.GetAsArray("handlers")
	handlers := []ConfigHandler{}
	//
	// fmt.Printf("GetConfigHandlers: entries: %v\n", entries)

	return handlers
}
