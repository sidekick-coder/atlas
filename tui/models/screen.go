package models

import (
	tea "charm.land/bubbletea/v2"
	"charm.land/bubbles/v2/key"
	"github.com/sidekick-coder/atlas/internal/app"
)

type ScreenPayload struct {
	App *app.App
	Options map[string]any
	ScreenIndex int
	Program *tea.Program
}

// Screen is the interface every top-level screen must implement.
type Screen interface {
	Title() string
	Init() tea.Cmd
	SetSize(width, height int)
	Update(msg tea.Msg) tea.Cmd
	GetBindings() []key.Binding
	Render() string
}


type ScreenFactory func(payload ScreenPayload) (Screen, error)
