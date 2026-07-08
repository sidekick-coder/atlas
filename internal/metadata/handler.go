package metadata

import (
	// "fmt"

	"github.com/sidekick-coder/atlas/internal/config"
	"github.com/sidekick-coder/atlas/internal/metadata/json"
	"github.com/sidekick-coder/atlas/internal/models"
)

type Handler interface {
	ID() string
	Extract(info *models.EntryInfo) (map[string]string, error)
	Set(info *models.EntryInfo, name string, value string) (bool, error)
	Unset(info *models.EntryInfo, name string) error
}

func GetHandlers(info *models.EntryInfo) []Handler {
	handlers := []Handler{}

	if (info.Type == "directory") {
		// return handlers
	}

	if (info.Ext == ".md") {
		handlers = append(handlers, MarkdownHandler{})
	}

	return handlers
}

func (m *Meta) GetHandlers() []Handler {
	handlers := []Handler{}

	if m.info.Type == "directory" {
		// return handlers
	}

	if m.info.Ext == ".md" {
		handlers = append(handlers, MarkdownHandler{})
	}

	if m.info.Ext == ".json" {
		handlers = append(handlers, json.Handler{})
	}

	return handlers
}

func (m *Meta) SetHandlersFromConfig(config *config.Config) {
	// handlers := config.GetConfigHandlers()
	//
	// fmt.Printf("Config handlers: %v\n", handlers)
}
