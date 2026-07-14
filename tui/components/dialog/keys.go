
package dialog

import (
	tea "charm.land/bubbletea/v2"

	key "github.com/sidekick-coder/atlas/tui/features/key"
)

type Keymap struct {
	Close key.Binding
}

var tags = []string{"component:dialog"}

var Bindings = Keymap{
	Close: key.CreateBinding("<esc>", "q").
		SetTags(tags...).
		SetDescription("close").
		SetHelp("esc"),
}

func (c *Component) GetBindings() []key.Binding {
	return []key.Binding{
		Bindings.Close,
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
	if !c.open {
		return nil
	}

	if (key.Matches(Bindings.Close)) {
		c.Close()
	}

	return nil
}


