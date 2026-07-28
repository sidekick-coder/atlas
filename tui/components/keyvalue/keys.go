package keyvalue

import (
	tea "charm.land/bubbletea/v2"
	"github.com/sidekick-coder/atlas/tui/features/key"
)

type Keymap struct {
	Up         key.Binding
	Down       key.Binding
	ScrollUp   key.Binding
	ScrollDown key.Binding
	JumpTop    key.Binding
	JumpBottom key.Binding
}

var tags = []string{"component:keyvalue"}

var Bindings = Keymap{
	Up: key.CreateBinding("k", "<up>").
		SetDescription("up").
		SetTags(tags...).
		SetHelp("k/up"),
	Down: key.CreateBinding("j", "<down>").
		SetDescription("down").
		SetTags(tags...).
		SetHelp("j/down"),
	ScrollUp: key.CreateBinding("p").
		SetDescription("scroll up").
		SetHidden(true).
		SetTags(tags...).
		SetHelp("<pageup>"),
	ScrollDown: key.CreateBinding("n").
		SetDescription("scroll down").
		SetHidden(true).
		SetTags(tags...).
		SetHelp("<pagedown>"),
	JumpTop: key.CreateBinding("<c-p>").
		SetDescription("jump to top").
		SetHidden(true).
		SetTags(tags...).
		SetHelp("<c-p>"),
	JumpBottom: key.CreateBinding("<c-n>").
		SetDescription("jump to bottom").
		SetHidden(true).
		SetTags(tags...).
		SetHelp("<c-n>"),
}

func (c *Component) GetBindigs() []key.Binding {
	return []key.Binding{
		Bindings.Up,
		Bindings.Down,
		Bindings.ScrollUp,
		Bindings.ScrollDown,
		Bindings.JumpTop,
		Bindings.JumpBottom,
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
		c.selection.Prev()

		cursor := c.selection.GetCursor()

		first := c.viewport.GetOffsetY()

		if first == cursor {
			c.viewport.Up()
		}

	}

	if key.Matches(Bindings.Down) {
		c.selection.Next()

		cursor := c.selection.GetCursor()

		last := c.viewport.GetOffsetY() + c.height

		if last == cursor {
			c.viewport.Down()
		}

	}

	if key.Matches(Bindings.ScrollUp) {
		c.viewport.Up()
	}

	if key.Matches(Bindings.ScrollDown) {
		c.viewport.Down()
	}

	if key.Matches(Bindings.JumpTop) {
		cursor := max(0, c.selection.GetCursor()-c.height)

		c.viewport.SetOffsetY(cursor)
		c.selection.SetCursor(cursor)
	}

	if key.Matches(Bindings.JumpBottom) {
		cursor := min(c.selection.GetTotal()-1, c.selection.GetCursor()+c.height)

		c.viewport.SetOffsetY(cursor)
		c.selection.SetCursor(cursor)
	}

	return nil
}
