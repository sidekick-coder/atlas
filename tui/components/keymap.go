package components

import "github.com/charmbracelet/bubbles/key"

// KeyMap holds all key bindings used across the TUI.
type KeyMap struct {
	Up          key.Binding
	Down        key.Binding
	PageNext    key.Binding
	PagePrev    key.Binding
	FocusNext   key.Binding
	MetaReplace key.Binding
	MetaUpdate  key.Binding
	MetaEditor  key.Binding
	MetaAdd     key.Binding
	SyncEntry   key.Binding
	Help        key.Binding
	Quit        key.Binding
}

// DefaultKeyMap is the default set of key bindings.
var DefaultKeyMap = KeyMap{
	Up: key.NewBinding(
		key.WithKeys("k", "up"),
		key.WithHelp("↑/k", "up"),
	),
	Down: key.NewBinding(
		key.WithKeys("j", "down"),
		key.WithHelp("↓/j", "down"),
	),
	PageNext: key.NewBinding(
		key.WithKeys("l", "right", "ctrl+f"),
		key.WithHelp("l/→", "next page"),
	),
	PagePrev: key.NewBinding(
		key.WithKeys("h", "left", "ctrl+b"),
		key.WithHelp("h/←", "prev page"),
	),
	FocusNext: key.NewBinding(
		key.WithKeys("tab"),
		key.WithHelp("tab", "switch panel"),
	),
	MetaReplace: key.NewBinding(
		key.WithKeys("r"),
		key.WithHelp("r", "replace"),
	),
	MetaUpdate: key.NewBinding(
		key.WithKeys("u"),
		key.WithHelp("u", "update"),
	),
	MetaEditor: key.NewBinding(
		key.WithKeys("e"),
		key.WithHelp("e", "editor"),
	),
	MetaAdd: key.NewBinding(
		key.WithKeys("a"),
		key.WithHelp("a", "add meta"),
	),
	SyncEntry: key.NewBinding(
		key.WithKeys("s"),
		key.WithHelp("s", "sync entry"),
	),
	Help: key.NewBinding(
		key.WithKeys("?"),
		key.WithHelp("?", "help"),
	),
	Quit: key.NewBinding(
		key.WithKeys("q", "ctrl+c"),
		key.WithHelp("q", "quit"),
	),
}


