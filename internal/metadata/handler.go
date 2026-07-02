package metadata

import (
	"github.com/sidekick-coder/atlas/internal/drive/v2"
)

type Handler interface {
	ID() string
	Extract(info *drive.EntryInfo) (map[string]string, error)
	Set(info *drive.EntryInfo, name string, value string) error
}
