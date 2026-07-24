package entrymeta

import (
	tea "charm.land/bubbletea/v2"
	"github.com/sidekick-coder/atlas/tui/components/toast"
	"github.com/sidekick-coder/atlas/tui/features/entrycontroller"
	"github.com/sidekick-coder/atlas/tui/features/key"
	"github.com/sidekick-coder/atlas/tui/messages"
)

type Keymap struct {
	Up     key.Binding
	Down   key.Binding
	Edit   key.Binding
	Add    key.Binding
	Delete key.Binding

	Sync   key.Binding
	Reload key.Binding
}

var Bindings = Keymap{
	Up: key.CreateBinding("k", "up").
		SetDescription("up").
		SetHelp("k/up"),
	Down: key.CreateBinding("j", "down").
		SetDescription("down").
		SetHelp("j/down"),
	Edit: key.CreateBinding("e").
		SetDescription("edit").
		SetHelp("e"),
	Add: key.CreateBinding("a", "add").
		SetDescription("add").
		SetHelp("a"),
	Delete: key.CreateBinding("d", "delete").
		SetDescription("delete").
		SetHelp("d"),
	Sync: key.CreateBinding("s").
		SetDescription("sync").
		SetHelp("s"),
	Reload: key.CreateBinding("r").
		SetDescription("reload").
		SetHelp("r"),
}

func (c *Component) GetBindigs() []key.Binding {
	return []key.Binding{
		Bindings.Up,
		Bindings.Down,
		Bindings.Edit,
		Bindings.Add,
		Bindings.Delete,
		Bindings.Reload,
		Bindings.Sync,
	}
}

func (c *Component) LoadBindings() tea.Cmd {
	key.Register(c.GetBindigs()...)
	return nil
}

func (c *Component) UnloadBindings() tea.Cmd {
	key.Unregister(c.GetBindigs()...)
	return nil
}

func (c *Component) HandleBindings(msg tea.KeyMsg) tea.Cmd {
	if key.Matches(Bindings.Up) {
		c.MoveUp()
	}

	if key.Matches(Bindings.Down) {
		c.MoveDown()
	}

	if key.Matches(Bindings.Edit) {

		selected, exists := c.GetSelected()

		if exists {
			c.dialog.SetTitle("Edit")
			c.dialog.SetContent(selected.Value)
			c.dialog.Open()

			c.dialog.OnSubmit(func(input string) tea.Cmd {
				return entrycontroller.Set(c.path, selected.Name, input)
			})

			return messages.SkipCmd()
		}
	}

	if key.Matches(Bindings.Add) {
		c.dialog.SetTitle("Meta name")
		c.dialog.SetContent("")
		c.dialog.Open()

		c.dialog.OnSubmit(func(input string) tea.Cmd {
			return entrycontroller.Set(c.path, input, "")
		})

		return messages.SkipCmd()
	}

	if key.Matches(Bindings.Delete) {
		selected, exists := c.GetSelected()

		if exists {
			return entrycontroller.Unset(c.path, selected.Name)
		}
	}

	if key.Matches(Bindings.Reload) {
		c.Load()

		return messages.ToastSuccessCmd("Reloaded")
	}

	if key.Matches(Bindings.Sync) {
		return tea.Batch(
			entrycontroller.Sync(c.path),
			toast.Success("Synced"),
		)
	}

	return nil
}
