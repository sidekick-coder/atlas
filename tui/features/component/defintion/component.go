package defintion

import tea "charm.land/bubbletea/v2"

type Component interface {
	SetProps(map[string]any)

	Init() tea.Cmd
	Update(msg tea.Msg) tea.Cmd
	Dispose() tea.Cmd

	Activate() tea.Cmd
	Deactivate() tea.Cmd

	Resize(width, height int)
	Render() string
}
