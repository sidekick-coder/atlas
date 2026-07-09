package root

import (
	"charm.land/bubbles/v2/key"
	tea "charm.land/bubbletea/v2"
)


func (m *model) GetBindings() []key.Binding {
	bindings := []key.Binding{}


	bindings = append(bindings, m.GetGlobalBindings()...)
	bindings = append(bindings, m.GetUserBindings()...)
	bindings = append(bindings, m.GetScreenManagerBindings()...)

	s, ok := m.GetCurrentScreen()

	if ok {
		bindings = append(bindings, s.GetBindings()...)
	}

	return bindings
}

func (m *model) LoadBindings() tea.Cmd {
	bindings := m.GetBindings()

	m.footer.SetBindings(bindings...)

	return nil
}

