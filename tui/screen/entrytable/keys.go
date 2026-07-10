package entrytable

import (
	tea "charm.land/bubbletea/v2"

	bkey "charm.land/bubbles/v2/key"
	tkey "github.com/sidekick-coder/atlas/tui/features/key"
)

type Keymap struct {
	Next tkey.Binding
	Prev tkey.Binding
}

var tags = []string{"entry_table"}

var Bindings = Keymap{
	Next: tkey.CreateBinding("n", "l").
		SetTags(tags...).
		SetDescription("Next page").
		SetHelp("n/l"),
	Prev: tkey.CreateBinding("p", "h").
		SetTags(tags...).
		SetDescription("Previous page").
		SetHelp("p/h"),
}

func (s *Screen) GetBindings() []bkey.Binding {
	bindings := []bkey.Binding{}

	return bindings
}

func (s *Screen) Bindings() []tkey.Binding {
	return []tkey.Binding{
		Bindings.Next,
		Bindings.Prev,
	}
}

func (s *Screen) LoadBindings() tea.Cmd {
	tkey.Register(s.Bindings()...)
	return nil
}

func (s *Screen) UnloadBindings() tea.Cmd {
	tkey.Unregister(s.Bindings()...)
	return nil
}

func (s *Screen) HadleBinding(msg tea.KeyMsg) tea.Cmd {
	if (tkey.Matches(Bindings.Next)) {
		s.loader.Next()
		s.loader.Load()
	}

	if (tkey.Matches(Bindings.Prev)) {
		s.loader.Prev()
		s.loader.Load()
	}
	return nil
}
