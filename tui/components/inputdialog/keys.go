package inputdialog

import (
	tea "charm.land/bubbletea/v2"

	key "github.com/sidekick-coder/atlas/tui/features/key"
)

type Keymap struct {
	Submit key.Binding
}

var tags = []string{"component:dialog"}

var Bindings = Keymap{
	Submit: key.CreateBinding("<enter>").
		SetTags(tags...).
		SetDescription("submit").
		SetHelp("enter"),
}

func (c *Component) GetBindings() []key.Binding {
	return []key.Binding{
		Bindings.Submit,
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
	if key.Matches(Bindings.Submit) {
		return c.submit()
	}

	return nil
}


