package metadata

import (
	"fmt"
	"maps"
	// "strings"
	"github.com/sidekick-coder/atlas/internal/models"
)

func (m *Meta) ExtractMap() (map[string]string, error) {
	result := make(map[string]string)

	for _, handler := range m.handlers {
		data, err := handler.Extract(m.info)

		if err != nil {
			return nil, fmt.Errorf("failed to extract metadata from handler %s(%s): %w", handler.GetID(), handler.GetTypeID(), err)
		}

		maps.Copy(result, data)
	}

	// result["basename"] = m.info.BaseName
	// result["type"] = m.info.Type
	// result["ext"] = strings.TrimPrefix(m.info.Ext, ".")

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

