package form

import (
	tea "charm.land/bubbletea/v2"
	"github.com/sidekick-coder/atlas/tui/features/key"
)

type Keymap struct {
	Up    key.Binding
	Down  key.Binding
	Enter key.Binding
	Cancel  key.Binding
}

var tags = []string{"component:form"}

var Binding = Keymap{
	Up: key.CreateBinding("<shift+tab>", "<Up>").
		SetHelp("shift+tab").
		SetDescription("Move up").
		SetTags(tags...),
	Down: key.CreateBinding("<tab>", "<Down>").
		SetHelp("tab").
		SetTags(tags...).
		SetDescription("Move down"),
	Enter: key.CreateBinding("<Enter>").
		SetHelp("enter").
		SetTags(tags...).
		SetDescription("Submit"),
	Cancel: key.CreateBinding("<esc>").
		SetHelp("esc").
		SetTags(tags...).
		SetDescription("Close form"),
}

func (c *Component) GetBindings() []key.Binding {
	return []key.Binding{
		Binding.Up,
		Binding.Down,
		Binding.Enter,
		Binding.Cancel,
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
	if !c.active {
		return nil
	}

	if key.Matches(Binding.Up) {
		c.focus.Prev()
	}

	if key.Matches(Binding.Down) {
		c.focus.Next()
	}

	if key.Matches(Binding.Cancel) {
		c.Events.Cancel.Emit()
		return nil
	}

	if key.Matches(Binding.Enter) {
		return c.submit()
	}

	index := c.focus.GetIndex()

	if index < 0 || index >= len(c.fields) {
		return nil
	}

	return nil
}
