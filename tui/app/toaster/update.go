package toaster

import (
	"log/slog"
	"time"

	tea "charm.land/bubbletea/v2"
	"github.com/sidekick-coder/atlas/tui/messages"
)

func (c *Component) Update(msg tea.Msg) tea.Cmd {
	if _, ok := msg.(tea.WindowSizeMsg); ok {
		c.LoadPosition()
	}

	if tm, ok := msg.(messages.Toast); ok {
		slog.Info("toaster messagele received", "title", tm.Title, "message", tm.Message, "color", tm.Color, "seconds", tm.Seconds)

		c.LoadPosition()

		c.toast.SetTitle(tm.Title)
		c.toast.SetContent(tm.Message)
		c.toast.SetColor(tm.Color)

		c.toast.SetActive(true)

		return func() tea.Msg {
			time.Sleep(time.Duration(tm.Seconds) * time.Second)

			return ToastExpiredMsg{}
		}
	}

	if _, ok := msg.(ToastExpiredMsg); ok {
		c.toast.SetActive(false)
	}

	return nil
}
