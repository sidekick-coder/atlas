package messages

import (
	tea "charm.land/bubbletea/v2"
)

type Action struct {
	Name string
	Context map[string]any
}

func ActionCmd(name string) tea.Cmd {
    return func() tea.Msg {
        return Action{
            Name: name,
			Context: map[string]any{},
        }
    }
}

func ActionWithContextCmd(name string, context map[string]any) tea.Cmd {
	return func() tea.Msg {
		return Action{
			Name:    name,
			Context: context,
		}
	}
}

