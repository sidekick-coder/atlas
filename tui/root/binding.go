package root

import (
	"charm.land/bubbles/v2/key"
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

func (m *model) LoadBindings() {
	bindings := m.GetBindings()

	m.footer.SetBindings(bindings...)
}

