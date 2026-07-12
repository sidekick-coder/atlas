package component

import tea "charm.land/bubbletea/v2"

var Definitions = map[string]DefinitionFactory{}

type Definition interface {
	Render() string
	Update(msg tea.Msg) tea.Cmd
	Init() tea.Cmd
	Dispose() tea.Cmd
	OnFocus()
	OnBlur()
}

type DefinitionFactory func(props ...map[string]any) (Definition, error)

func RegisterDefinition[T Definition](name string, factory func(props ...map[string]any) (T, error)) {
	Definitions[name] = func(props ...map[string]any) (Definition, error) {
		return factory(props...)
	}
}
