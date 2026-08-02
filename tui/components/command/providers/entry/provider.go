package entry

import (
	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
	"github.com/sidekick-coder/atlas/internal/models"
	"github.com/sidekick-coder/atlas/internal/repository/entry"
	"github.com/sidekick-coder/atlas/tui/app/program"
	"github.com/sidekick-coder/atlas/tui/app/screen"
	"github.com/sidekick-coder/atlas/tui/components/command/provider"
	"github.com/sidekick-coder/atlas/tui/features/theme"
	"github.com/sidekick-coder/atlas/tui/screen/entrysingle"
)

type Provider struct {
	color      lipgloss.Style
	commands   []provider.Command
	repository *entry.Repository
}

func Create() *Provider {
	p := &Provider{
		color:      theme.BaseStyle().Foreground(theme.Accent()),
		commands:   []provider.Command{},
		repository: program.GetApp().EntryRepo(),
	}

	return p
}

func (p *Provider) AddCommand(name string, description string, execute func() tea.Cmd) {
	p.commands = append(p.commands, provider.Command{
		Name:        name,
		Description: description,
		Execute:     execute,
	})
}

func (p *Provider) AddEntryCommand(entry models.Entry) {
	p.AddCommand(p.color.Render("[entry]")+" (open) open "+entry.Path, "open entry page metas", func() tea.Cmd {
		options := map[string]any{}
		options["entry"] = entry
		options["path"] = entry.Path
		return screen.Add(entrysingle.ID, options)
	})

}

func (p *Provider) List(payload provider.ListPayload) []provider.Command {
	p.commands = []provider.Command{}

	q := payload.Query

	if q == "" {
		return p.commands
	}

	opts := entry.ListOptions{
		Query: []string{q},
		Limit: 10,
	}

	entries, err := p.repository.List(opts)

	if err != nil {
		return p.commands
	}

	for _, e := range entries {
		p.AddEntryCommand(e)
	}

	return p.commands
}
