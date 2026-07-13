package form

import (
	tea "charm.land/bubbletea/v2"
	"github.com/sidekick-coder/atlas/tui/features/key"
	"github.com/sidekick-coder/atlas/tui/messages"
)

type Keymap struct {
	Up    key.Binding
	Down  key.Binding
	Enter key.Binding
	Blur key.Binding
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
	Blur: key.CreateBinding("<esc>").
		SetHelp("esc").
		SetTags(tags...).
		SetDescription("Close form"),
}

func (c *Component) GetBindings() []key.Binding {
	return []key.Binding{
		Binding.Up,
		Binding.Down,
		Binding.Enter,
		Binding.Blur,
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
	if !c.focused {
		return nil
	}

	if key.Matches(Binding.Up) {
		c.selection.Prev()
		c.Refresh()
		return messages.SkipCmd()
	}

	if key.Matches(Binding.Down) {
		c.selection.Next()
		c.Refresh()
		return messages.SkipCmd()
	}

	if key.Matches(Binding.Blur) {
		c.selection.SetCursor(-1)
		c.Refresh()
		return messages.SkipCmd()
	}

	if key.Matches(Binding.Enter) {
		return c.submit()
	}

	return nil
}
