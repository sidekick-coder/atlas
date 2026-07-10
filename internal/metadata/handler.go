package metadata

import (
	"fmt"

	"github.com/sidekick-coder/atlas/internal/config"
	"github.com/sidekick-coder/atlas/internal/fs"
	"github.com/sidekick-coder/atlas/internal/metadata/content"
	"github.com/sidekick-coder/atlas/internal/metadata/frontmatter"
	"github.com/sidekick-coder/atlas/internal/metadata/json"
	"github.com/sidekick-coder/atlas/internal/metadata/stat"
	"github.com/sidekick-coder/atlas/internal/models"
)

func (m *Meta) SetHandlersFromConfig(config *config.Config) error {
	configHandlers, err := config.GetConfigHandlers()
	handlers := []models.MetaHandler{}

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

		payload := models.MetaHandlerPayload{
			ID:      hc.ID,
			Options: hc.Options,
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

		return fmt.Errorf("unknown handler type: %s", hc.Type)
	}

	m.handlers = handlers

	return nil
}
