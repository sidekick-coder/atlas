package components

import "github.com/charmbracelet/bubbles/key"

// KeyMap holds all key bindings used across the TUI.
type KeyMap struct {
	Up          key.Binding
	Down        key.Binding
	FocusNext   key.Binding
	MetaReplace key.Binding
	MetaUpdate  key.Binding
	MetaEditor  key.Binding
	MetaAdd     key.Binding
	Help        key.Binding
	Quit        key.Binding
}

// DefaultKeyMap is the default set of key bindings.
var DefaultKeyMap = KeyMap{
	Up: key.NewBinding(
		key.WithKeys("k", "up"),
		key.WithHelp("↑/k", "move up"),
	),
	Down: key.NewBinding(
		key.WithKeys("j", "down"),
		key.WithHelp("↓/j", "move down"),
	),
	FocusNext: key.NewBinding(
		key.WithKeys("tab"),
		key.WithHelp("tab", "switch panel"),
	),
	MetaReplace: key.NewBinding(
		key.WithKeys("r"),
		key.WithHelp("r", "replace value"),
	),
	MetaUpdate: key.NewBinding(
		key.WithKeys("u"),
		key.WithHelp("u", "update value"),
	),
	MetaEditor: key.NewBinding(
		key.WithKeys("e"),
		key.WithHelp("e", "edit in editor"),
	),
	MetaAdd: key.NewBinding(
		key.WithKeys("a"),
		key.WithHelp("a", "add meta"),
	),
	Help: key.NewBinding(
		key.WithKeys("?"),
		key.WithHelp("?", "show help"),
	),
	Quit: key.NewBinding(
		key.WithKeys("q", "ctrl+c"),
		key.WithHelp("q", "quit"),
	),
}


