package screen

import (
	tea "charm.land/bubbletea/v2"
	"github.com/charmbracelet/bubbles/key"
)

// Screen is the interface every top-level screen must implement.
type Screen interface {
	Title() string
	Init() tea.Cmd
	SetSize(width, height int)
	Update(msg tea.Msg) tea.Cmd
	GetBindings() []key.Binding
	Render() string
}
