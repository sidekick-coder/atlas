package list

import (
	tea "charm.land/bubbletea/v2"
	"github.com/sidekick-coder/atlas/tui/features/key"
	"github.com/sidekick-coder/atlas/tui/messages"
)

type Keymap struct {
	Up     key.Binding
	Down   key.Binding
	Next   key.Binding
	Prev   key.Binding
	Select key.Binding
}

var tags = []string{"component:list"}

var Binding = Keymap{
	Up: key.CreateBinding("k", "<up>").
		SetTags(tags...).
		SetHelp("k").
		SetDescription("up"),
	Down: key.CreateBinding("j").
		SetTags(tags...).
		SetHelp("j").
		SetDescription("down"),
	Select: key.CreateBinding("<enter>").
		SetTags(tags...).
		SetHelp("<enter>").
		SetDescription("Select item"),
}

func (c *Component) GetBindigs() []key.Binding {
	return []key.Binding{
		Binding.Up,
		Binding.Down,
		Binding.Select,
	}
}

func (c *Component) LoadBindings() {
	key.Register(c.GetBindigs()...)
}

func (c *Component) UnloadBindings() {
	key.Unregister(c.GetBindigs()...)
}

func (c *Component) HandleBinding(km tea.KeyMsg) tea.Cmd {
	if key.Matches(Binding.Up) {
		return c.Up()
	}

	if key.Matches(Binding.Down) {
		return c.Down()
	}

	if key.Matches(Binding.Select) {
		index, ok := c.GetCursor()

		if ok && c.onSelect != nil {
			return c.onSelect(index)
		}

		return messages.SkipCmd()
	}

	return nil
}
