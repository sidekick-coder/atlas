package empty

import (
	tea "charm.land/bubbletea/v2"
	"github.com/sidekick-coder/atlas/tui/features/key"
)

type Keymap struct {
	Enter key.Binding
}

var tags = []string{"screen:empty"}

var Binding = Keymap{
	Enter: key.CreateBinding("<enter>").
		SetDescription("select").
		SetHelp("enter").
		SetTags(tags...),
}

func (s *Screen) GetBindings() []key.Binding {
	return []key.Binding{
		Binding.Enter,
	}
}

func (s *Screen) LoadBindings() {
	key.Register(s.GetBindings()...)
}

func (s *Screen) UnloadBindings() {
	key.Unregister(s.GetBindings()...)
}

func (s *Screen) HandleBinding(km tea.KeyMsg) tea.Cmd {
	if key.Matches(Binding.Enter) {
		return s.Select(s.selection.GetCursor())
	}

	return nil
}
