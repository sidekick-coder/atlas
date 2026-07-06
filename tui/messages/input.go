
package messages

import tea "charm.land/bubbletea/v2"

type Input struct {
	Title    string 
	InitialValue  string
	Callback func(value string) tea.Cmd
}

func InputCmd(msg Input) tea.Cmd {
	return func() tea.Msg {
		return msg
	}
}
