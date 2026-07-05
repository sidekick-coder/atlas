
package messages

import tea "charm.land/bubbletea/v2"

type Input struct {
	Title    string 
	InitialValue  string
	Callback func(value string) tea.Cmd
}

func InputCmd(title string, callback func(value string) tea.Cmd) tea.Cmd {
	return func() tea.Msg {
		return Input{
			Title:    title,
			InitialValue:  "",
			Callback: callback,
		}
	}
}

func InputWithInitialCmd(title string, initial string, callback func(value string) tea.Cmd) tea.Cmd {
	return func() tea.Msg {
		return Input{
			Title:    title,
			InitialValue:  initial,
			Callback: callback,
		}
	}
}
