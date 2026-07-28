package defintion

import (
	tea "charm.land/bubbletea/v2"
	"github.com/sidekick-coder/atlas/tui/features/context"
)

type Component interface {
	SetProps(map[string]any)

	Context() *context.Feature

	Init() tea.Cmd
	Update(msg tea.Msg) tea.Cmd
	Dispose() tea.Cmd

	Activate() tea.Cmd
	Deactivate() tea.Cmd

	Resize(width, height int)
	Render() string
}
