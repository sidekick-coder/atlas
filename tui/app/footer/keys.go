package footer

import (
	"log/slog"

	tea "charm.land/bubbletea/v2"
	"github.com/sidekick-coder/atlas/tui/features/key"
)

type Keymap struct {
	Toggle key.Binding
}

var tags = []string{"global"}

var Binding = Keymap{
	Toggle: key.CreateBinding("?", "<f1>").
		SetHelp("?").
		SetDescription("help").
		SetTags(tags...),
}

func (c *Component) GetBindings() []key.Binding {
	return []key.Binding{
		Binding.Toggle,
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
	if key.Matches(Binding.Toggle) {
		if c.dialog.IsOpen() {
			c.dialog.Close()
			slog.Info("Closing help dialog")
			return nil
		}

		slog.Info("Opening help dialog")

		c.dialog.Open()
	}

	return nil
}
