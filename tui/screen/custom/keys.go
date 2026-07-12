package custom

import (
	tea "charm.land/bubbletea/v2"
	"github.com/sidekick-coder/atlas/tui/features/key"
	"github.com/sidekick-coder/atlas/tui/messages"
)

type KeyMap struct {
	Up   key.Binding
	Down key.Binding
	Esc  key.Binding
}

var tags = []string{"screen:custom"}

var Bindings = KeyMap{
	Esc: key.CreateBinding("<esc>").
		SetTags(tags...).
		SetDescription("unfocus components").
		SetHelp("esc"),
	Up: key.CreateBinding("<leader>k", "<leader><up>p").
		SetTags(tags...).
		SetDescription("<leader>k").
		SetHelp("<leader>k"),
	Down: key.CreateBinding("<leader>j", "<leader><down>", "<leader>n").
		SetTags(tags...).
		SetDescription("down").
		SetHelp("<leader>j"),
}

func (s *Screen) GetBindings() []key.Binding {
	return []key.Binding{
		Bindings.Up,
		Bindings.Down,
		Bindings.Esc,
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

	if key.Matches(Bindings.Esc) {
		s.Select(-1)
		return messages.SkipCmd()
	}

	return nil
}
