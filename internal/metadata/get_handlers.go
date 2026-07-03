package metadata 

import (
	"github.com/sidekick-coder/atlas/internal/models"
)

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
