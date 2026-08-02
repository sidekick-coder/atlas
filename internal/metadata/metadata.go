package metadata

import (
	"github.com/sidekick-coder/atlas/internal/metadata/handler"
	"github.com/sidekick-coder/atlas/internal/models"
)

type Meta struct {
	info     *models.EntryInfo
	handlers []handler.Handler
}

func Handler(info *models.EntryInfo) (*Meta, error) {
	meta := &Meta{
		info:     info,
		handlers: []handler.Handler{},
	}

	return meta, nil
}

// Deprecated: Create is deprecated, use Handler instead. It will be removed in future versions.
func Create(info *models.EntryInfo) (*Meta, error) {
	return Handler(info)
}

func (m *Meta) Set(name string, value string) (bool, error) {
	handlers := m.handlers
	info := m.info
	success := false

	for _, handler := range handlers {
		updated, err := handler.Set(info, name, value)

		if err != nil {
			return success, err
		}

		if updated {
			success = true
		}
	}

	return success, nil
}

func (m *Meta) Unset(name string) (bool, error) {
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
