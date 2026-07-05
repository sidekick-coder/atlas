package root

import (
	"github.com/charmbracelet/bubbles/key"
)

type GlobalKeyMap struct {
	Quit      key.Binding
	OpenEntry key.Binding

	NextScreen key.Binding
	PrevScreen key.Binding

	SyncAll key.Binding
}

var GlobalBindings = GlobalKeyMap{
	Quit: key.NewBinding(
		key.WithKeys("q", "ctrl+c"),
		key.WithHelp("q", "quit"),
	),
	OpenEntry: key.NewBinding(
		key.WithKeys("e"),
		key.WithHelp("e", "open entry screen"),
	),
	NextScreen: key.NewBinding(
		key.WithKeys("tab"),
		key.WithHelp("tab", "next screen"),
	),
	PrevScreen: key.NewBinding(
		key.WithKeys("shift+tab"),
		key.WithHelp("shift+tab", "previous screen"),
	),
	SyncAll: key.NewBinding(
		key.WithKeys("S"),
		key.WithHelp("S", "sync all entries"),
	),
}

func (m *model) GetBindings() []key.Binding {
	bindings := []key.Binding{}

	s := m.screens[m.currentIndex]

	bindings = append(bindings, s.GetBindings()...)
	bindings = append(bindings, GlobalBindings.Quit)
	bindings = append(bindings, GlobalBindings.OpenEntry)
	bindings = append(bindings, GlobalBindings.NextScreen)
	bindings = append(bindings, GlobalBindings.PrevScreen)
	bindings = append(bindings, GlobalBindings.SyncAll)

	actions := m.app.Config().GetKeymapsByGroup("global")

	for _, action := range actions {
		bindings = append(bindings, key.NewBinding(
			key.WithKeys(action.Keys...),
			key.WithHelp(action.Keys[0], action.Description),
		))
	}

	return bindings
}
