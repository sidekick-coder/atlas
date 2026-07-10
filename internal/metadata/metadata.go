package metadata

import (
	"fmt"
	"slices"

	"github.com/sidekick-coder/atlas/internal/metadata/handler"
	"github.com/sidekick-coder/atlas/internal/models"
)

type Meta struct {
	info *models.EntryInfo
	handlers []handler.Handler
}

var SytemMetaNames = []string{
	"basename",
	"type",
	"ext",
}

func Handler(info *models.EntryInfo) (*Meta, error) {
	meta := &Meta{
		info: info,
		handlers: []handler.Handler{},
	}

	return meta, nil
}

// Deprecated: Create is deprecated, use Handler instead. It will be removed in future versions.
func Create(info *models.EntryInfo) (*Meta, error) {
	return Handler(info)
}


func (m *Meta) Set(name string, value string) (bool, error) {
	if slices.Contains(SytemMetaNames, name) {
		return false, fmt.Errorf("cannot set system meta: %s", name)
	}

	handlers := m.handlers
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

func (m *Meta) Unset(name string) (bool, error) {
	if slices.Contains(SytemMetaNames, name) {
		return false, fmt.Errorf("cannot unset system meta: %s", name)
	}

	handlers := m.handlers
	info := m.info

	for _, handler := range handlers {
		err := handler.Unset(info, name)

		if err != nil {
			return false, err
		}
	}

	return true, nil
}

