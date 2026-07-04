package metadata

import (
	"github.com/sidekick-coder/atlas/internal/models"
)

type Handler interface {
	ID() string
	Extract(info *models.EntryInfo) (map[string]string, error)
	Set(info *models.EntryInfo, name string, value string) (bool, error)
	Unset(info *models.EntryInfo, name string) error
}
