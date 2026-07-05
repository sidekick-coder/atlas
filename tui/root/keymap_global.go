package root

import (
	tea "charm.land/bubbletea/v2"
	"github.com/charmbracelet/bubbles/key"
)

type GlobalKeyMap struct {
	Quit      key.Binding
	OpenEntry key.Binding

	NextScreen key.Binding
	PrevScreen key.Binding
}

var GlobalBindings = GlobalKeyMap{
	Quit: key.NewBinding(
		key.WithKeys("q", "ctrl+c"),
		key.WithHelp("q", "quit"),
	),
	OpenEntry: key.NewBinding(
		key.WithKeys("e"),
		key.WithHelp("e", "entries"),
	),
	NextScreen: key.NewBinding(
		key.WithKeys("tab"),
		key.WithHelp("tab", "next screen"),
	),
	PrevScreen: key.NewBinding(
		key.WithKeys("shift+tab"),
		key.WithHelp("shift+tab", "previous screen"),
	),
}

func (m *model) GetGlobalBindings() []key.Binding {
	bindings := []key.Binding{}

	bindings = append(bindings, GlobalBindings.Quit)
	bindings = append(bindings, GlobalBindings.OpenEntry)
	bindings = append(bindings, GlobalBindings.NextScreen)
	bindings = append(bindings, GlobalBindings.PrevScreen)

	return bindings
}

func (m *model) HandleGlobalKeyMap(msg tea.Msg) tea.Cmd {
	keyMsg, ok := msg.(tea.KeyMsg)

	if !ok {
		return nil
	}

	if key.Matches(keyMsg, GlobalBindings.Quit) {
		return tea.Quit
	}

	if key.Matches(keyMsg, GlobalBindings.OpenEntry) {
		m.AddScreen("entry", nil)
		return  nil
	}

	if key.Matches(keyMsg, GlobalBindings.NextScreen) {
		nextIndex := (m.currentIndex + 1) % len(m.screens)
		m.currentIndex = nextIndex
		m.tabBar.SetCurrent(nextIndex)
		return nil
	}

	if key.Matches(keyMsg, GlobalBindings.PrevScreen) {
		prevIndex := (m.currentIndex - 1 + len(m.screens)) % len(m.screens)
		m.currentIndex = prevIndex
		m.tabBar.SetCurrent(prevIndex)

		return nil
	}

	return nil
}
