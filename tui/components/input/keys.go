package input

import (
	tea "charm.land/bubbletea/v2"
	"github.com/sidekick-coder/atlas/tui/features/key"
	"github.com/sidekick-coder/atlas/tui/messages"
)

type Keymap struct {
	Left  key.Binding
	Right key.Binding
	Enter key.Binding
	Close key.Binding
}

var tags = []string{"component:input"}

var Binding = Keymap{
	Left: key.CreateBinding("<Left>", "<c-h>", "<c-p").
		SetHelp("/c-h/c-p").
		SetTags(tags...).
		SetHidden(true).
		SetDescription("Move cursor left"),
	Right: key.CreateBinding("<Right>", "<c-l>", "<c-n>").
		SetHelp("/c-l/c-n").
		SetHidden(true).
		SetDescription("Move cursor right").
		SetTags(tags...),
}

func (c *Component) GetBindings() []key.Binding {
	return []key.Binding{
		Binding.Left,
		Binding.Right,
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

func (i *Input) HandleKeypress(msg tea.Msg) tea.Cmd {
	if !i.enabled {
		return nil
	}

	textMsg, ok := msg.(tea.KeyPressMsg)

	if !ok {
		return nil
	}

	if key.Matches(Binding.Left) {
		if i.cursor > 0 {
			i.cursor--
		}

		return messages.SkipCmd()
	}

	if key.Matches(Binding.Right) {
		if i.cursor < len(i.buf) {
			i.cursor++
		}

		return messages.SkipCmd()
	}

	code := textMsg.Code

	if code == tea.KeyBackspace {
		if i.cursor > 0 {
			i.buf = append(i.buf[:i.cursor-1], i.buf[i.cursor:]...)
			i.cursor--
		}

		return messages.SkipCmd()
	}

	if code == tea.KeyDelete {
		if i.cursor < len(i.buf) {
			i.buf = append(i.buf[:i.cursor], i.buf[i.cursor+1:]...)
		}

		return messages.SkipCmd()
	}

	if code == tea.KeyHome {
		i.cursor = 0

		return messages.SkipCmd()
	}

	if code == tea.KeyEnd {
		i.cursor = len(i.buf)

		return messages.SkipCmd()
	}

	if textMsg.Text != "" {
		i.buf = append(i.buf[:i.cursor], append([]rune(textMsg.Text), i.buf[i.cursor:]...)...)
		i.cursor += len([]rune(textMsg.Text))

		return messages.SkipCmd()
	}

	return nil

}
