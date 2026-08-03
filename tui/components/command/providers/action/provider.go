package action

import (
	"slices"
	"strings"

	tea "charm.land/bubbletea/v2"
	"github.com/sidekick-coder/atlas/tui/action"
	"github.com/sidekick-coder/atlas/tui/app/program"
	"github.com/sidekick-coder/atlas/tui/components/command/provider"
	"github.com/sidekick-coder/atlas/tui/features/theme"
)

type Provider struct {
	commands []provider.Command
}

func Create() *Provider {
	p := &Provider{
		commands: []provider.Command{},
	}

	p.Load()

	return p
}

func (p *Provider) AddCommand(name string, description string, execute func() tea.Cmd) {
	p.commands = append(p.commands, provider.Command{
		Name:        name,
		Description: description,
		Execute:     execute,
	})
}

func (p *Provider) Load() {
	color := theme.BaseStyle().Foreground(theme.Error())

	p.AddCommand(color.Render("[action] ")+"sync all", "sync all entries", func() tea.Cmd {
		ctx := make(map[string]any)

		return action.Execute("entry-sync-all", ctx)
	})

	manager := program.GetApp().Action

	actions, err := manager.List()

	if err != nil {
		return
	}

	for _, a := range actions {
		if a.Options["command"] == nil || a.Options["command"] == false {
			continue
		}

		p.AddCommand(color.Render("[action] ")+a.ID, a.Type, func() tea.Cmd {
			ctx := make(map[string]any)

			return action.Execute(a.ID, ctx)
		})
	}

	slices.SortFunc(p.commands, func(a, b provider.Command) int {
		return strings.Compare(a.Name, b.Name)
	})

}

func (p *Provider) List(payload provider.ListPayload) []provider.Command {
	commands := []provider.Command{}

	q := payload.Query

	if q == "" {
		return p.commands
	}

	for _, entry := range p.commands {
		if entry.MatchQuery(q) {
			commands = append(commands, entry)
		}
	}

	return commands
}
