package toast

import (
	"log/slog"

	tea "charm.land/bubbletea/v2"
	"github.com/sidekick-coder/atlas/tui/app/program"
	"github.com/sidekick-coder/atlas/tui/messages"
)

func Error(msg string, timeout ...int) tea.Cmd {
	slog.Error("toast error", "message", msg)
	program.Send(messages.ToastErrorMessage(msg, timeout...))
	return nil
	// return messages.ToastErrorCmd(msg, timeout...)
}

func Success(msg string, timeout ...int) tea.Cmd {
	return messages.ToastSuccessCmd(msg, timeout...)
}
