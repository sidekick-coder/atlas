package columnlist

import (
	tea "charm.land/bubbletea/v2"
	"github.com/sidekick-coder/atlas/tui/components/table/column"
	"github.com/sidekick-coder/atlas/tui/features/key"
	"github.com/sidekick-coder/atlas/tui/messages"
)

type Keymap struct {
	Up    key.Binding
	Down  key.Binding
	Enter key.Binding
	Add   key.Binding
	Remove key.Binding
	Close key.Binding
}

var tags = []string{"columnlist", "column list"}

var Binding = Keymap{
	Up: key.CreateBinding("k", "<Up>").
		SetHelp("k/up").
		SetTags(tags...).
		SetDescription("up"),
	Down: key.CreateBinding("j", "<Down>").
		SetHelp("j/down").
		SetTags(tags...).
		SetDescription("down"),
	Enter: key.CreateBinding("<enter>").
		SetHelp("enter").
		SetTags(tags...).
		SetDescription("select"),
	Add: key.CreateBinding("a").
		SetHelp("a").
		SetTags(tags...).
		SetDescription("add"),
	Remove: key.CreateBinding("d").
		SetHelp("d").
		SetTags(tags...).
		SetDescription("remove"),
	Close: key.CreateBinding("<esc>", "q").
		SetHelp("esc").
		SetTags(tags...).
		SetDescription("close"),
}

func (c *Component) GetBindings() []key.Binding {
	return []key.Binding{
		Binding.Up,
		Binding.Down,
		Binding.Enter,
		Binding.Add,
		Binding.Remove,
		Binding.Close,
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
	if !c.IsOpen() {
		return nil
	}

	if key.Matches(Binding.Enter) {
		return c.EditCurrent()
	}

	if key.Matches(Binding.Up) {
		c.column.Selection.Prev()
	}

	if key.Matches(Binding.Down) {
		c.column.Selection.Next()
	}

	if key.Matches(Binding.Close) {
		c.Close()
	}

	if key.Matches(Binding.Add) {
		c.column.AddColumn(column.Column{
			Label: "New Column",
			Field: "new_column",
			Width: 0,
		})
	}

	if key.Matches(Binding.Remove) {
		if len(c.column.GetColumns()) == 1 {
			return messages.ToastErrorCmd("Cannot remove the last column")
		}

		c.column.RemoveSelectedColumn()
	}

	return nil
}
