
package messages

import tea "charm.land/bubbletea/v2"

type Input struct {
	Title    string 
	Callback func(value string) tea.Cmd
}

func InputCmd(title string, callback func(value string) tea.Cmd) tea.Cmd {
	return func() tea.Msg {
		return Input{
			Title:    title,
			Callback: callback,
		}
	}
}
