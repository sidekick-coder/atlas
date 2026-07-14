package table

import (
	tea "charm.land/bubbletea/v2"
	"github.com/sidekick-coder/atlas/tui/features/key"
)

type Keymap struct {
	Up          key.Binding
	Down        key.Binding
	Enter       key.Binding
	EditColumns key.Binding
}

var tags = []string{"component:table"}

var Binding = Keymap{
	Up: key.CreateBinding("k", "<Up>").
		SetHelp("k").
		SetTags(tags...).
		SetDescription("up"),
	Down: key.CreateBinding("j", "<Down>").
		SetHelp("j").
		SetTags(tags...).
		SetDescription("down"),
	Enter: key.CreateBinding("<enter>").
		SetHelp("enter").
		SetTags(tags...).
		SetDescription("select"),
	EditColumns: key.CreateBinding("ec").
		SetHelp("ec").
		SetTags(tags...).
		SetDescription("columns"),
}

func (c *Component) GetBindings() []key.Binding {
	return []key.Binding{
		Binding.Up,
		Binding.Down,
		Binding.Enter,
		Binding.EditColumns,
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
	if key.Matches(Binding.Up) {
		c.itemSelection.Prev()
	}

	if key.Matches(Binding.Down) {
		c.itemSelection.Next()
	}

	if key.Matches(Binding.EditColumns) {
		c.columnList.Open()
	}

	if key.Matches(Binding.Enter) {
		if c.onSelect != nil {
			return c.onSelect(c.itemSelection.GetCursor())
		}
	}

	return nil
}
