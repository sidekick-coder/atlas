package entrylist

import (
	tea "charm.land/bubbletea/v2"

	key "github.com/sidekick-coder/atlas/tui/features/key"
)

type Keymap struct {
	Next key.Binding
	Prev key.Binding
}

var tags = []string{"screen:entry_table"}

var Bindings = Keymap{
	Next: key.CreateBinding("<tab>").
		SetTags(tags...).
		SetDescription("focus next").
		SetHelp("tab"),
	Prev: key.CreateBinding("<shift+tab>").
		SetTags(tags...).
		SetDescription("focus prev").
		SetHelp("shift+tab"),
}

func (s *Screen) GetBindings() []key.Binding {
	return []key.Binding{
		Bindings.Next,
		Bindings.Prev,
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

func (s *Screen) HadleBinding(msg tea.KeyMsg) tea.Cmd {
	if key.Matches(Bindings.Next) {
		return s.focus.Next()
	}

	if key.Matches(Bindings.Prev) {
		return s.focus.Prev()
	}

	return nil
}
