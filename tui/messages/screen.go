package messages

import tea "charm.land/bubbletea/v2"

type AddScreen struct {
	Name string 
	Options map[string]any
}

type RemoveScreen struct {
	Index int
}

type ReplaceCurrentScreen struct {
	Index int
	Name  string
	Options map[string]any
}


func AddScreenCmd(msg AddScreen) tea.Cmd {
	return func() tea.Msg {
		return msg
	}
}

func RemoveScreenCmd(msg RemoveScreen) tea.Cmd {
	return func() tea.Msg {
		return msg
	}
}

func ReplaceScreenCmd(msg ReplaceCurrentScreen) tea.Cmd {
	return func() tea.Msg {
		return msg
	}
}

