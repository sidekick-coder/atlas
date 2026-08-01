package toast

import (

	tea "charm.land/bubbletea/v2"
	"github.com/sidekick-coder/atlas/tui/app/program"
	"github.com/sidekick-coder/atlas/tui/messages"
)

func Error(msg string, timeout ...int) tea.Cmd {
	program.Send(messages.ToastErrorMessage(msg, timeout...))
	return nil
}

func Success(msg string, timeout ...int) tea.Cmd {
	return messages.ToastSuccessCmd(msg, timeout...)
}
