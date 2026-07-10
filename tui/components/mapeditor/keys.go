package mapeditor

import (
	"log"

	tea "charm.land/bubbletea/v2"
	"github.com/sidekick-coder/atlas/tui/features/key"
	"github.com/sidekick-coder/atlas/tui/messages"
)

type Keymap struct {
	Up    key.Binding
	Down  key.Binding
	Close key.Binding
}

var tags = []string{"mapeditor", "map editor"}

var Binding = Keymap{
	Up: key.CreateBinding("<shift+tab>", "<Up>").
		SetHelp("k/up").
		SetDescription("Move up").
		SetTags(tags...),
	Down: key.CreateBinding("<tab>", "<Down>").
		SetHelp("j/down").
		SetTags(tags...).
		SetDescription("Move down"),
	Close: key.CreateBinding("<Esc>").
		SetHelp("esc").
		SetTags(tags...).
		SetDescription("Close map editor"),
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
		log.Println("Up key pressed")
		c.selection.Prev()
	}

	if key.Matches(Binding.Down) {
		c.selection.Next()
	}

	if key.Matches(Binding.Close) {
		c.Close()
	}

	return messages.SkipCmd()
}
