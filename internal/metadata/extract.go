package metadata

import (
	"fmt"
	"maps"

	// "strings"
	"github.com/sidekick-coder/atlas/internal/metadata/handler"
	"github.com/sidekick-coder/atlas/internal/models"
)

func (m *Meta) ExtractMap() (map[string]string, error) {
	result := make(map[string]string)

	for _, h := range m.handlers {
		payload := handler.ExtractPayload{
			Info:  m.info,
			Metas: result,
		}

		data, err := h.Extract(payload)

		if err != nil {
			return nil, fmt.Errorf("failed to extract metadata from handler %s(%s): %w", h.GetID(), h.GetTypeID(), err)
		}

		maps.Copy(result, data)
	}

	return result, nil
}

func (m *Meta) Extract() ([]models.EntryMeta, error) {
	metas, err := m.ExtractMap()

	if err != nil {
		return nil, err
	}

	result := []models.EntryMeta{}

	for key, value := range metas {
		meta := models.EntryMeta{
			Name:  key,
			Value: value,
		}

		result = append(result, meta)
	}

	return result, nil
}
