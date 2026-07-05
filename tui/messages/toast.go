package messages

import tea "charm.land/bubbletea/v2"

type Toast struct {
	Message string
	Level   string
	Duration int // Duration in seconds
}

func ToastCmd(mgs Toast) tea.Cmd {
	return func() tea.Msg {
		return mgs
	}
}

func ToastErrorCmd(message string, duration ...int) tea.Cmd {
	return ToastCmd(Toast{
		Message:  message,
		Level:    "error",
		Duration: duration[0],
	})
}

func ToastInfoCmd(message string, duration ...int) tea.Cmd {
	return ToastCmd(Toast{
		Message:  message,
		Level:    "info",
		Duration: duration[0],
	})
}

func ToastSuccessCmd(message string, duration ...int) tea.Cmd {
	return ToastCmd(Toast{
		Message:  message,
		Level:    "success",
		Duration: duration[0],
	})
}
