package entrytable

import (
	tea "charm.land/bubbletea/v2"

	tkey "charm.land/bubbles/v2/key"
)

type Keymap struct {
}

var Bindings = Keymap{
}

func (s *Screen) GetBindings() []tkey.Binding {
	bindings := []tkey.Binding{}

	return bindings
}

func (s *Screen) RegisterBindings() tea.Cmd {
	return nil
}

func (s *Screen) HandleKeypress(msg tea.KeyMsg) tea.Cmd {
	return nil
}
