package component

import tea "charm.land/bubbletea/v2"

var Definitions = map[string]func(DefinitionPayload) Definition{}

type Definition interface {
	Render() string
	Update(msg tea.Msg) tea.Cmd
	Init() tea.Cmd
	Dispose() tea.Cmd
	OnFocus()
	OnBlur()
}

type DefinitionPayload struct {
	Width   int
	Height  int
	Options map[string]any
}

type DefinitionFactory func(payload DefinitionPayload) Definition

func RegisterDefinition(name string, factory DefinitionFactory) {
	Definitions[name] = factory
}
