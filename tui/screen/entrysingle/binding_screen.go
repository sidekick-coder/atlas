package entrysingle

import (
	"charm.land/bubbles/v2/key"
)

type ScreenKeyMap struct {
	Up         key.Binding
	Down       key.Binding
	Edit        key.Binding
	Reload     key.Binding
	Sync	   key.Binding
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
	Edit: key.NewBinding(
		key.WithKeys("e"),
		key.WithHelp("e", "edit value"),
	),
	Sync: key.NewBinding(
		key.WithKeys("s"),
		key.WithHelp("s", "sync"),
	),
	Reload: key.NewBinding(
		key.WithKeys("r"),
		key.WithHelp("r", "reload"),
	),
}

func (s *Screen) GetScreenBindigs() []key.Binding {
	return []key.Binding{
		ScreenBindings.Up,
		ScreenBindings.Down,
		ScreenBindings.Edit,
		ScreenBindings.Reload,
	}
}
