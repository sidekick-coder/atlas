package list

import (
	tea "charm.land/bubbletea/v2"
	// "github.com/sidekick-coder/atlas/tui/app/program"
	"github.com/sidekick-coder/atlas/tui/features/key"
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
}

func (c *Component) GetBindigs() []key.Binding {
	return []key.Binding{
		Binding.Up,
		Binding.Down,
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
		c.selection.Prev()

		// return tea.Batch(
		// 	program.Command(UpMsg{}),
		// 	program.Command(MovedMsg{}),
		// )
	}

	if key.Matches(Binding.Down) {
		c.selection.Next()

		// return tea.Batch(
		// 	program.Command(DownMsg{}),
		// 	program.Command(MovedMsg{}),
		// )
	}

	return nil
}
