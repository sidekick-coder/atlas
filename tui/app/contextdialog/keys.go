package contextdialog

import (
	tea "charm.land/bubbletea/v2"
	"github.com/sidekick-coder/atlas/tui/features/key"
)

type Keymap struct {
	Open key.Binding
}

var tags = []string{"global"}

var Binding = Keymap{
	Open: key.CreateBinding("<f2>").
		SetDescription("context").
		SetTags(tags...).
		SetHelp("f2"),
}

func (c *Component) GetBindings() []key.Binding {
	return []key.Binding{
		Binding.Open,
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
	if c.dialog.IsOpen() {
		return c.kv.Update(msg)
	}

	if !c.dialog.IsOpen() && key.Matches(Binding.Open) {
		c.dialog.Open()
		c.Load()
		c.kv.Activate()
		return nil
	}

	if c.dialog.IsOpen() && key.Matches(Binding.Open) {
		c.dialog.Close()
		c.kv.Deactivate()
		return nil
	}

	return nil
}
