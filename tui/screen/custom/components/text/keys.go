package text 

import (
	tea "charm.land/bubbletea/v2"

	"github.com/sidekick-coder/atlas/tui/features/key"
)

type Keymap struct {
	Up key.Binding
	Down key.Binding
}

var tags = []string{"component:text"}

var Bindings = Keymap{
	Up: key.CreateBinding("k", "<up>").
		SetTags(tags...).
		SetDescription("Scroll up").
		SetHelp("k/up"),
	Down: key.CreateBinding("j", "<down>").
		SetTags(tags...).
		SetDescription("Scroll down").
		SetHelp("j/down"),
}

func (c *Component) GetBindings() []key.Binding {
	return []key.Binding{
		Bindings.Up,
		Bindings.Down,
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

func (c *Component) HandleBinding(msg tea.KeyMsg) tea.Cmd {
	if key.Matches(Bindings.Up) {
		c.viewport.Up()
	}

	if key.Matches(Bindings.Down) {
		c.viewport.Down()
	}

	return nil
}
