package command

import (
	tea "charm.land/bubbletea/v2"

	key "github.com/sidekick-coder/atlas/tui/features/key"
)

type Keymap struct {
	Toggle key.Binding
	Up     key.Binding
	Down   key.Binding
	Enter  key.Binding
}

var tags = []string{"command"}

var Bindings = Keymap{
	Toggle: key.CreateBinding("<ctrl+p>").
		SetTags(tags...).
		SetDescription("open").
		SetHelp("ctrl+p"),
	Up: key.CreateBinding("<shift+tab>", "<up>", "<ctrl+k>").
		SetHelp("shift+tab").
		SetDescription("Move up").
		SetTags(tags...),
	Down: key.CreateBinding("<tab>", "<down>", "<ctrl+j>").
		SetHelp("tab").
		SetTags(tags...).
		SetDescription("Move down"),
	Enter: key.CreateBinding("<Enter>").
		SetHelp("enter").
		SetTags(tags...).
		SetDescription("Submit"),
}

func (c *Component) LoadBindings() tea.Cmd {
	key.Register(Bindings.Toggle)
	return nil
}

func (c *Component) LoadControlBindings() tea.Cmd {
	key.Register(Bindings.Up, Bindings.Down, Bindings.Enter)
	return nil
}

func (c *Component) UnloadBindings() tea.Cmd {
	key.Unregister(Bindings.Toggle)
	return nil
}

func (c *Component) UnloadControlBindings() tea.Cmd {
	key.Unregister(Bindings.Up, Bindings.Down, Bindings.Enter)
	return nil
}

func (c *Component) HadleBinding(msg tea.KeyMsg) tea.Cmd {
	if key.Matches(Bindings.Toggle) {
		c.dialog.Open()
	}

	if !c.dialog.IsOpen() {
		return nil
	}

	if key.Matches(Bindings.Up) {
		c.focus.Prev()
	}

	if key.Matches(Bindings.Down) {
		c.focus.Next()
	}

	if key.Matches(Bindings.Enter) {
		return c.Execute()
	}

	return nil
}
