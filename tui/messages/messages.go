package messages

import (
	tea "charm.land/bubbletea/v2"
)

type AddScreen struct {
	Name string
	Options map[string]any
}

func AddScreenCmd(name string, opts map[string]any) tea.Cmd {
    return func() tea.Msg {
        return AddScreen{
            Name: name,
            Options: opts,
        }
    }
}

