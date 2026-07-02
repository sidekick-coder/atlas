package metadata 

import (
	"github.com/sidekick-coder/atlas/internal/drive/v2"
)

func GetHandlers(info *drive.EntryInfo) []Handler {
	handlers := []Handler{}

	if (info.Type == "directory") {
		// return handlers
	}

	if (info.Ext == ".md") {
		handlers = append(handlers, MarkdownHandler{})
	}

	return handlers
}
