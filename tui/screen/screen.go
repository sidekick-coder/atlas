package screen

import (
	tea "charm.land/bubbletea/v2"
	"github.com/charmbracelet/bubbles/key"
)

// Screen is the interface every top-level screen must implement.
type Screen interface {
	// Title is shown in the tab bar.
	Title() string
	// Init is called once when the screen first becomes active.
	Init() tea.Cmd
	// SetSize is called when the screen size changes.
	SetSize(width, height int)
	// Update handles messages and returns commands.
	Update(msg tea.Msg) tea.Cmd
	// Render returns the screen content as a plain string (no tea.View).
	GetBindings() []key.Binding
	// The root model composites overlays on top before creating the final View.
	Render() string
}
