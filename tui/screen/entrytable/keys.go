package entrytable

import (
	tea "charm.land/bubbletea/v2"

	tkey "charm.land/bubbles/v2/key"
	"github.com/sidekick-coder/atlas/tui/features/key"
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
	key.Register()
	return nil
}

func (s *Screen) HandleKeypress(msg tea.Msg) tea.Cmd {
	return nil
}
