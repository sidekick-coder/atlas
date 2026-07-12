package toast

import tea "charm.land/bubbletea/v2"
import "github.com/sidekick-coder/atlas/tui/messages"

func Error(msg string, timeout ...int) tea.Cmd {
	return messages.ToastErrorCmd(msg, timeout...)
}

func Success(msg string, timeout ...int) tea.Cmd {
	return messages.ToastSuccessCmd(msg, timeout...)
}
