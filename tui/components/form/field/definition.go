package field

import tea "charm.land/bubbletea/v2"

type Definition interface {
	Render() string
	Resize(width, height int)

	SetValue(value string)
	GetValue() string

	SetProps(props map[string]any)

	Activate() tea.Cmd 
	Deactivate() tea.Cmd
	Update(msg tea.Msg) tea.Cmd
}

type DefinitionFunc func() Definition
