package screen

import (
	"slices"
	"strings"

	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
	"github.com/sidekick-coder/atlas/tui/app/screen"
	"github.com/sidekick-coder/atlas/tui/components/command/provider"
	"github.com/sidekick-coder/atlas/tui/features/theme"
)

type Entry struct {
	Title       string
	Description string
}

type Provider struct {
	entries  []screen.ScreenEntry
	commands []provider.Command
	color    lipgloss.Style
}

func Create() *Provider {
	p := &Provider{
		entries:  []screen.ScreenEntry{},
		commands: []provider.Command{},
		color:    theme.BaseStyle().Foreground(theme.Primary()),
	}

	entries, err := screen.List()

	if err == nil {
		p.entries = entries
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

func (p *Provider) AddScreenCommands(entry screen.ScreenEntry) {
	p.AddCommand(p.color.Render("[screen]")+" (open) open "+entry.ID, "open screen: "+entry.ID, func() tea.Cmd {
		return screen.Add(entry.ID, entry.Options)
	})

	p.AddCommand(p.color.Render("[screen]")+" (replace) replace "+entry.ID, "replace current screen: "+entry.ID, func() tea.Cmd {
		return screen.ReplaceCurrent(entry.ID, entry.Options)
	})
}

func (p *Provider) Load() {
	for _, entry := range p.entries {
		p.AddScreenCommands(entry)
	}

	p.AddCommand(p.color.Render("[screen]")+" close screen", "close current screen", func() tea.Cmd {
		return screen.RemoveCurrent()
	})

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
		if strings.Contains(strings.ToLower(entry.Name), strings.ToLower(q)) {
			commands = append(commands, entry)
		}
	}

	return commands
}
