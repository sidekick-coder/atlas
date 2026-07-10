package table

import (
	tea "charm.land/bubbletea/v2"
	"github.com/sidekick-coder/atlas/tui/features/key"
)

type Keymap struct {
	Up    key.Binding
	Down  key.Binding
	Enter key.Binding
	Right  key.Binding
	Left  key.Binding
}

var Binding = Keymap{
	Up:    key.CreateBinding("k", "up").SetHelp("k/up").SetDescription("Move up"),
	Down:  key.CreateBinding("j", "down").SetHelp("j/down").SetDescription("Move down"),
	Enter: key.CreateBinding("enter", "enter").SetHelp("enter").SetDescription("Select item"),
	Right:  key.CreateBinding("l", "next").SetHelp("l/next").SetDescription("Next column"),
	Left:  key.CreateBinding("h", "prev").SetHelp("h/prev").SetDescription("Previous column"),
}

func (c *Component) LoadBindings() tea.Cmd {
	key.Register(
		Binding.Up,
		Binding.Down,
		Binding.Enter,
		Binding.Right,
		Binding.Left,
	)

	return nil
}

func (c *Component) HandleBindings(msg tea.Msg) tea.Cmd {
	if key.Matches(Binding.Left) {
		c.columnSelection.Prev()
	}

	if key.Matches(Binding.Right) {
		c.columnSelection.Next()
	}

	if key.Matches(Binding.Up) {
		c.itemSelection.Prev()
	}

	if key.Matches(Binding.Down) {
		c.itemSelection.Next()
	}

	return nil
}
