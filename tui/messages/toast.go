package messages

import (
	tea "charm.land/bubbletea/v2"
	"github.com/sidekick-coder/atlas/tui/features/theme"
)

type Toast struct {
	Title   string
	Message string
	Color   string
	Seconds int // Duration in seconds
}

func ToastCmd(mgs Toast) tea.Cmd {
	return func() tea.Msg {
		return mgs
	}
}

func ToastErrorMessage(message string, seconds ...int) Toast {
	sec := 5
	color := theme.Current.Error
	title := "Error"

	if len(seconds) > 0 {
		sec = seconds[0]
	}

	return Toast{
		Title:   title,
		Color:   color,
		Message: message,
		Seconds: sec,
	}
}

func ToastSuccessMessage(message string, seconds ...int) Toast {
	sec := 5
	color := theme.Current.Success
	title := "Success"

	if len(seconds) > 0 {
		sec = seconds[0]
	}

	return Toast{
		Title:   title,
		Color:   color,
		Message: message,
		Seconds: sec,
	}
}

func ToastErrorCmd(message string, seconds ...int) tea.Cmd {
	return ToastCmd(ToastErrorMessage(message, seconds...))
}

func ToastSuccessCmd(message string, seconds ...int) tea.Cmd {
	return ToastCmd(ToastSuccessMessage(message, seconds...))
}
