package contextdialog

import (
	tea "charm.land/bubbletea/v2"
	"github.com/sidekick-coder/atlas/tui/features/key"
)

type Keymap struct {
	Toggle key.Binding
}

var tags = []string{"global"}

var Binding = Keymap{
	Toggle: key.CreateBinding("<f2>").
		SetDescription("context").
		SetTags(tags...).
		SetHelp("f2"),
}

func (c *Component) GetBindings() []key.Binding {
	return []key.Binding{
		Binding.Toggle,
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

func (c *Component) HandleBindings(msg tea.KeyMsg) tea.Cmd {
	if key.Matches(Binding.Toggle) {
		c.dialog.Toggle()
		return nil
	}

	if c.dialog.IsOpen() {
		return c.kv.Update(msg)
	}

	return nil
}
