package root

import (
	"github.com/charmbracelet/bubbles/key"
)


func (m *model) GetBindings() []key.Binding {
	bindings := []key.Binding{}

	s, ok := m.GetCurrentScreen()

	if ok {
		bindings = append(bindings, s.GetBindings()...)
	}

	bindings = append(bindings, m.GetGlobalBindings()...)
	bindings = append(bindings, m.GetUserBindings()...)
	bindings = append(bindings, m.GetScreenManagerBindings()...)

	return bindings
}
