package table

import tea "charm.land/bubbletea/v2"

type ChangeMsg struct {
	Index int
}

func SelectionChangeCmd(index int) tea.Cmd {
	return func() tea.Msg {
		return ChangeMsg{Index: index}
	}
}
