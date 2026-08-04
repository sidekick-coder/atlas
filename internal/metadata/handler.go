package metadata

import (
	"fmt"
	"slices"
	"strings"

	"github.com/sidekick-coder/atlas/internal/config"
	"github.com/sidekick-coder/atlas/internal/fs"
	"github.com/sidekick-coder/atlas/internal/metadata/handler"
	"github.com/sidekick-coder/atlas/internal/metadata/handlers/content"
	"github.com/sidekick-coder/atlas/internal/metadata/handlers/filemeta"
	"github.com/sidekick-coder/atlas/internal/metadata/handlers/frontmatter"
	"github.com/sidekick-coder/atlas/internal/metadata/handlers/json"
	"github.com/sidekick-coder/atlas/internal/metadata/handlers/setter"
	"github.com/sidekick-coder/atlas/internal/metadata/handlers/shell"
	"github.com/sidekick-coder/atlas/internal/metadata/handlers/stat"
)


func (m *Meta) SetHandlersFromConfig(config *config.Config) error {
	configHandlers, err := config.GetConfigHandlers()
	handlers := []handler.Handler{}

	if err != nil {
		return err
	}

	for _, hc := range configHandlers {
		if len(hc.Patterns) == 0 {
			return fmt.Errorf("handler %s has no patterns", hc.ID)
		}

		matched, err := fs.MatchAny(m.info.Path, hc.Patterns)


		if err != nil {
			return err
		}

		if !matched {
			continue
		}

		payload := handler.Payload{
			ID:      hc.ID,
			Options: hc.Options,
			Config:  config,
		}

		if hc.Type == "frontmatter" {
			handlers = append(handlers, frontmatter.Create(payload))
			continue
		}

		if hc.Type == "json" {
			handlers = append(handlers, json.Create(payload))
			continue
		}

		if hc.Type == "content" {
			handlers = append(handlers, content.Create(payload))
			continue
		}

		if hc.Type == "stat" {
			handlers = append(handlers, stat.Create(payload))
			continue
		}

		if hc.Type == "shell" {
			handlers = append(handlers, shell.Create(payload))
			continue
		}

		if hc.Type == "setter" {
			handlers = append(handlers, setter.Create(payload))
			continue
		}

		if hc.Type == "file-meta" {
			handlers = append(handlers, filemeta.Create(payload))
			continue
		}

		return fmt.Errorf("unknown handler type: %s", hc.Type)
	}

	m.handlers = handlers

	// sort by id
	slices.SortFunc(m.handlers, func(a, b handler.Handler) int {
		return strings.Compare(a.GetID(), b.GetID())
	})

	return nil
}
