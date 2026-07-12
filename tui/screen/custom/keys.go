package custom

import (
	tea "charm.land/bubbletea/v2"
	"github.com/sidekick-coder/atlas/tui/features/key"
	"github.com/sidekick-coder/atlas/tui/messages"
)

type KeyMap struct {
	Up   key.Binding
	Down key.Binding
}

var tags = []string{"screen", "custom"}

var Bindings = KeyMap{
	Up: key.CreateBinding("<leader>k", "<leader><up>").
		SetTags(tags...).
		SetDescription("<leader>k").
		SetHelp("k/up"),
	Down: key.CreateBinding("<leader>j", "<leader><down>").
		SetTags(tags...).
		SetDescription("down").
		SetHelp("j/down"),
}

func (s *Screen) GetBindings() []key.Binding {
	return []key.Binding{
		Bindings.Up,
		Bindings.Down,
	}
}

func (s *Screen) LoadBindings() tea.Cmd {
	key.Register(s.GetBindings()...)
	return nil
}

func (s *Screen) UnloadBindings() tea.Cmd {
	key.Unregister(s.GetBindings()...)
	return nil
}

func (s *Screen) HandleBinding(msg tea.KeyMsg) tea.Cmd {
	if key.Matches(Bindings.Up) {
		s.Next()
		return messages.SkipCmd()
	}

	if key.Matches(Bindings.Down) {
		s.Prev()
		return messages.SkipCmd()
	}

	return nil
}
