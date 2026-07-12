package screen

import (
	tea "charm.land/bubbletea/v2"
	"github.com/sidekick-coder/atlas/tui/features/key"
)

type Keymap struct {
	Next  key.Binding
	Prev  key.Binding
	Add   key.Binding
	Close key.Binding
}

var Bindings = Keymap{
	Next: key.CreateBinding("<leader>n"),
	Prev: key.CreateBinding("<leader>p"),
}

func (f *Feature) GetBindings() []key.Binding {
	return []key.Binding{
		Bindings.Next,
		Bindings.Prev,
	}
}

func (f *Feature) LoadBindings() {
	key.Register(f.GetBindings()...)
}

func (f *Feature) UnloadBindings() {
	key.Unregister(f.GetBindings()...)
}

func (f *Feature) HandleBinding(km tea.KeyMsg) tea.Cmd {
	if key.Matches(Bindings.Next) {
		f.Next()
	}

	if key.Matches(Bindings.Prev) {
		f.Prev()
	}

	return nil
}
