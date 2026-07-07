package metadata

import (
	"fmt"
	"slices"
	"strings"

	"github.com/sidekick-coder/atlas/internal/metadata/json"
	"github.com/sidekick-coder/atlas/internal/models"
)

type Meta struct {
	info     *models.EntryInfo
}

var SytemMetaNames = []string{"basename", "type", "ext", "handlers"}

func Create(info *models.EntryInfo) (*Meta, error) {
	meta := &Meta{
		info: info,
	}

	return meta, nil
}

func (m *Meta) GetHandlers() []Handler {
	handlers := []Handler{}

	if (m.info.Type == "directory") {
		// return handlers
	}

	if (m.info.Ext == ".md") {
		handlers = append(handlers, MarkdownHandler{})
	}

	if (m.info.Ext == ".json") {
		handlers = append(handlers, json.Handler{})
	}

	return handlers
}

func (m *Meta) ExtractMap() (map[string]string, error) {
	handlers := m.GetHandlers()
	result := make(map[string]string)

	ids := []string{}

	for _, handler := range handlers {
		data, err := handler.Extract(m.info)

		ids = append(ids, handler.ID())

		if err != nil {
			fmt.Printf("Error extracting metadata with handler %s: %v\n", handler.ID(), err)
			continue
		}

		for key, value := range data {
			result[key] = value
		}
	}

	result["basename"] = m.info.BaseName
	result["type"] = m.info.Type
	result["ext"] = strings.TrimPrefix(m.info.Ext, ".")
	result["handlers"] = strings.Join(ids, ",")

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
			Name:   key,
			Value: value,
		}

		result = append(result, meta)
	}

	return result, nil
}


func (m *Meta) Set(name string, value string) (bool, error) {
	if slices.Contains(SytemMetaNames, name) {
		return false, fmt.Errorf("cannot set system field: %s", name)
	}

	handlers := m.GetHandlers()
	info := m.info

	for _, handler := range handlers {
		updated, err := handler.Set(info, name, value)

		if err != nil {
			return false, err
		}

		if updated {
			return true, nil
		}
	}

	return false, nil
}
