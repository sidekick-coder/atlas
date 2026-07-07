package messages

import tea "charm.land/bubbletea/v2"

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
	sec := 3
	color := "196"
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
	sec := 3
	color := "46"
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
