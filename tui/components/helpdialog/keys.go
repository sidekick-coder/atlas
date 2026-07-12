package helpdialog

import (
	tea "charm.land/bubbletea/v2"
	"github.com/sidekick-coder/atlas/tui/features/key"
	"github.com/sidekick-coder/atlas/tui/messages"
)

type Keymap struct {
	Up    key.Binding
	Down  key.Binding
	Close key.Binding
}

var tags = []string{"global:help"}

var Binding = Keymap{
	Up: key.CreateBinding("k", "<up>").
		SetHelp("k/up").
		SetDescription("Move up").
		SetTags(tags...),
	Down: key.CreateBinding("j", "<down>").
		SetHelp("j/down").
		SetTags(tags...).
		SetDescription("Move down"),
	Close: key.CreateBinding("<esc>", "q").
		SetHelp("esc/q").
		SetTags(tags...).
		SetDescription("Close help"),
}

func (c *Component) GetBindings() []key.Binding {
	return []key.Binding{
		Binding.Up,
		Binding.Down,
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

	if key.Matches(Binding.Up) {
		c.viewport.Up()
		return messages.SkipCmd()
	}

	if key.Matches(Binding.Down) {
		c.viewport.Down()
		return messages.SkipCmd()
	}

	if key.Matches(Binding.Close) {
		c.Close()
	}

	// trap all other key events to prevent them from propagating to the parent component
	return messages.SkipCmd()
}
