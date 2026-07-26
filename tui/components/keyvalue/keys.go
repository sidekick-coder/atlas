package keyvalue

import (
	tea "charm.land/bubbletea/v2"
	"github.com/sidekick-coder/atlas/tui/app/program"
	"github.com/sidekick-coder/atlas/tui/features/key"
)

type Keymap struct {
	Up     key.Binding
	Down   key.Binding
}

var Bindings = Keymap{
	Up: key.CreateBinding("k", "up").
		SetDescription("up").
		SetHelp("k/up"),
	Down: key.CreateBinding("j", "down").
		SetDescription("down").
		SetHelp("j/down"),
}

func (c *Component) GetBindigs() []key.Binding {
	return []key.Binding{
		Bindings.Up,
		Bindings.Down,
	}
}

func (c *Component) LoadBindings() tea.Cmd {
	key.Register(c.GetBindigs()...)
	return nil
}

func (c *Component) UnloadBindings() tea.Cmd {
	key.Unregister(c.GetBindigs()...)
	return nil
}

func (c *Component) HandleBindings(msg tea.KeyMsg) tea.Cmd {
	if key.Matches(Bindings.Up) {
		c.selection.Prev()

		return program.Command(UpMsg{})
	}

	if key.Matches(Bindings.Down) {
		c.selection.Next()

		return program.Command(DownMsg{})
	}

	return nil
}
