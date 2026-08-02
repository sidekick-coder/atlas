package provider

import tea "charm.land/bubbletea/v2"

type Command struct {
	Name        string
	Description string
	Execute     func() tea.Cmd
	Render      func() string
}

type ListPayload struct {
	Query string
}

type Provider interface {
	List(payload ListPayload) []Command
}

