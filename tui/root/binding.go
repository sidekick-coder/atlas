package root

import (
	"github.com/charmbracelet/bubbles/key"
)


func (m *model) GetBindings() []key.Binding {
	bindings := []key.Binding{}

	s := m.screens[m.currentIndex]

	bindings = append(bindings, s.GetBindings()...)

	bindings = append(bindings, m.GetGlobalBindings()...)
	bindings = append(bindings, m.GetUserBindings()...)

	return bindings
}
