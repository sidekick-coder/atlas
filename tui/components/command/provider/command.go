package provider

import (
	"strings"

	tea "charm.land/bubbletea/v2"
)

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

func (c *Command) MatchQuery(query string) bool {
	queryWords := strings.Fields(strings.ToLower(query))

	if len(queryWords) == 0 {
		return true
	}

	text := strings.ToLower(c.Name + " " + c.Description)

	for _, word := range queryWords {
		if !strings.Contains(text, word) {
			return false
		}
	}

	return true
}
