package handler

import (
	"github.com/sidekick-coder/atlas/internal/config"
	"github.com/sidekick-coder/atlas/internal/models"
)

type Payload struct {
	ID      string
	Options map[string]any
	Config  *config.Config
}

type Handler interface {
	GetID() string
	GetTypeID() string
	Extract(info *models.EntryInfo) (map[string]string, error)
	Set(info *models.EntryInfo, name string, value string) (bool, error)
	Unset(info *models.EntryInfo, name string) error
}
