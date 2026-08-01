package textfield

import (
	tea "charm.land/bubbletea/v2"

	key "github.com/sidekick-coder/atlas/tui/features/key"
)

type Keymap struct {
}

var tags = []string{"component:textfield"}

var Bindings = Keymap{
}

func (c *Component) GetBindings() []key.Binding {
	return []key.Binding{
	}
}

func (c *Component) LoadBindings() tea.Cmd {
	key.Register(c.GetBindings()...)
	return nil
}

func (c *Component) UnloadBindings() tea.Cmd {
	key.Unregister(c.GetBindings()...)
	return nil
}

func (c *Component) HadleBinding(msg tea.KeyMsg) tea.Cmd {
	return nil
}

