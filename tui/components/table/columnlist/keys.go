package columnlist

import (
	"log"

	tea "charm.land/bubbletea/v2"
	"github.com/sidekick-coder/atlas/tui/features/key"
)

type Keymap struct {
	Up    key.Binding
	Down  key.Binding
	Enter key.Binding
	Close  key.Binding
}

var Binding = Keymap{
	Up:    key.CreateBinding("k", "<Up>").SetHelp("k/up").SetDescription("Move up"),
	Down:  key.CreateBinding("j", "<Down>").SetHelp("j/down").SetDescription("Move down"),
	Enter: key.CreateBinding("enter").SetHelp("enter").SetDescription("Select item"),
	Close:  key.CreateBinding("<Esc>").SetHelp("esc").SetDescription("Close table"),
}

func (c *Component) LoadBindings() tea.Cmd {
	key.Register(
		Binding.Up,
		Binding.Down,
		Binding.Enter,
		Binding.Close,
	)

	return nil
}

func (c *Component) HandleBindings(msg tea.Msg) tea.Cmd {
	log.Println("Handling bindings in columnlist")
	if key.Matches(Binding.Up) {
		c.column.Selection.Prev()
	}

	if key.Matches(Binding.Down) {
		c.column.Selection.Next()
	}

	if key.Matches(Binding.Close) {
		c.Close()
	}

	return nil
}
