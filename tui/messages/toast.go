package messages

import tea "charm.land/bubbletea/v2"

type Toast struct {
	Message string
	Level   string
	Duration int // Duration in seconds
}

func ToastCmd(message string, level string, duration int) tea.Cmd {
	return func() tea.Msg {
		return Toast{
			Message: message,
			Level:   level,
			Duration: duration,
		}
	}
}

func ToastErrorCmd(message string, duration int) tea.Cmd {
	return ToastCmd(message, "error", duration)
}
