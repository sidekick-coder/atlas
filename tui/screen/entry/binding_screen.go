package entry

import (
	"charm.land/bubbles/v2/key"
)

type ScreenKeyMap struct {
	Up     key.Binding
	Down   key.Binding
	Enter  key.Binding
	Next   key.Binding
	Prev   key.Binding
	Sync   key.Binding
	Reload key.Binding
	Search key.Binding
}

var ScreenBindings = ScreenKeyMap{
	Up: key.NewBinding(
		key.WithKeys("k", "up"),
		key.WithHelp("k/up", "up"),
	),
	Down: key.NewBinding(
		key.WithKeys("j", "down"),
		key.WithHelp("j/down", "down"),
	),
	Next: key.NewBinding(
		key.WithKeys("l", "right"),
		key.WithHelp("l/right", "next"),
	),
	Prev: key.NewBinding(
		key.WithKeys("h", "left"),
		key.WithHelp("h/left", "prev"),
	),
	Enter: key.NewBinding(
		key.WithKeys("enter"),
		key.WithHelp("enter", "select"),
	),
	Sync: key.NewBinding(
		key.WithKeys("s"),
		key.WithHelp("s", "sync"),
	),
	Reload: key.NewBinding(
		key.WithKeys("r"),
		key.WithHelp("r", "reload"),
	),
	Search: key.NewBinding(
		key.WithKeys("/"),
		key.WithHelp("/", "search"),
	),
}

func (s *Screen) GetScreenBindigs() []key.Binding {
	bindings := []key.Binding{}

	bindings = append(bindings,
		ScreenBindings.Up,
		ScreenBindings.Down,
		ScreenBindings.Next,
		ScreenBindings.Prev,
		ScreenBindings.Sync,
		ScreenBindings.Enter,
		ScreenBindings.Reload,
		ScreenBindings.Search,
	)

	return bindings
}

