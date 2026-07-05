package input

import (
	tea "charm.land/bubbletea/v2"
	"github.com/sidekick-coder/atlas/tui/messages"
)

func (i *Input) HandleKeypress(msg tea.Msg) tea.Cmd {
	if !i.Active {
		return nil
	}

	textMsg, ok := msg.(tea.KeyPressMsg)

	if !ok {
		return nil
	}

	code := textMsg.Code

	if code == tea.KeyEscape {
		i.Close()
		return messages.SkipCmd()
	}

	if code == tea.KeyEnter {
		value := string(i.Buf)

		i.Close()

		if i.callback != nil {
			return i.callback(value)
		}

		return messages.SkipCmd()
	}

	if code == tea.KeyBackspace {
		if i.Cursor > 0 {
			i.Buf = append(i.Buf[:i.Cursor-1], i.Buf[i.Cursor:]...)
			i.Cursor--
		}
	}

	if code == tea.KeyDelete {
		if i.Cursor < len(i.Buf) {
			i.Buf = append(i.Buf[:i.Cursor], i.Buf[i.Cursor+1:]...)
		}
	}

	if code == tea.KeyLeft {
		if i.Cursor > 0 {
			i.Cursor--
		}
	}

	if code == tea.KeyRight {
		if i.Cursor < len(i.Buf) {
			i.Cursor++
		}
	}

	if code == tea.KeyHome {
		i.Cursor = 0
	}

	if code == tea.KeyEnd {
		i.Cursor = len(i.Buf)
	}

	if textMsg.Text != "" {
		i.Buf = append(i.Buf[:i.Cursor], append([]rune(textMsg.Text), i.Buf[i.Cursor:]...)...)
		i.Cursor += len([]rune(textMsg.Text))
	}

	return messages.SkipCmd()
}
