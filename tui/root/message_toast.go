package root

import (
	"time"

	tea "charm.land/bubbletea/v2"
	"github.com/sidekick-coder/atlas/tui/messages"
)

func (m *model) HandleToastMessages(msg tea.Msg) tea.Cmd {
	if t, ok := msg.(messages.Toast); ok {
		m.toaster.SetColor(t.Color)
		m.toaster.SetTitle(t.Title)
		m.toaster.SetContent(t.Message)
		m.toaster.SetActive(true)

		return func() tea.Msg {
			dur := time.Duration(t.Seconds) * time.Second

			time.Sleep(dur)

			m.toaster.SetActive(false)

			return messages.Skip{}
		}
	}

	return nil
}
