package messages

import tea "charm.land/bubbletea/v2"

type Skip struct {
}

func SkipCmd() tea.Cmd {
	return func() tea.Msg {
		return nil
	}
}
