package entrysingle

import (
	tea "charm.land/bubbletea/v2"
	"github.com/sidekick-coder/atlas/tui/features/key"
)

type Keymap struct {
}

var Bindings = Keymap{}

func (s *Screen) GetBindigs() []key.Binding {
	return []key.Binding{}
}

func (s *Screen) LoadBindings() tea.Cmd {
	key.Register(s.GetBindigs()...)
	return nil
}

func (s *Screen) UnloadBindings() tea.Cmd {
	key.Unregister(s.GetBindigs()...)
	return nil
}

func (s *Screen) HandleBindings(msg tea.KeyMsg) tea.Cmd {
	return nil
}
