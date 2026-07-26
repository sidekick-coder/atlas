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

type ExtractPayload struct {
	Info	*models.EntryInfo 
	Metas	map[string]string
}

type Handler interface {
	GetID() string
	GetTypeID() string
	Extract(ExtractPayload) (map[string]string, error)
	Set(info *models.EntryInfo, name string, value string) (bool, error)
	Unset(info *models.EntryInfo, name string) error
}
