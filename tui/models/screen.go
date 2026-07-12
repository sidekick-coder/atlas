package models

import (
	tea "charm.land/bubbletea/v2"
	"github.com/sidekick-coder/atlas/internal/app"
)

type ScreenPayload struct {
	App *app.App
	Options map[string]any
}

// Screen is the interface every top-level screen must implement.
type Screen interface {
	Title() string
	Init() tea.Cmd
	SetSize(width, height int)
	Update(msg tea.Msg) tea.Cmd
	Render() string
	Dispose() tea.Cmd
}


type ScreenFactory func(payload ScreenPayload) (Screen, error)
