package empty

import (
	tea "charm.land/bubbletea/v2"
	"github.com/sidekick-coder/atlas/tui/features/key"
)

type Keymap struct {
}

var Binding = Keymap{}

func (s *Screen) GetBindings() []key.Binding {
	return []key.Binding{
	}
}

func (s *Screen) LoadBindings() {
	key.Register(s.GetBindings()...)
}

func (s *Screen) UnloadBindings() {
	key.Unregister(s.GetBindings()...)
}

func (s *Screen) HandleBinding(km tea.KeyMsg) tea.Cmd {
	return nil
}
