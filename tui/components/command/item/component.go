package item

import (
	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
	"github.com/charmbracelet/x/ansi"
	"github.com/sidekick-coder/atlas/internal/logger"
	"github.com/sidekick-coder/atlas/tui/components/command/provider"
	"github.com/sidekick-coder/atlas/tui/features/theme"
)

type Component struct {
	cmd    provider.Command
	width  int
	height int
	active bool
}

func Create(cmd provider.Command) *Component {
	return &Component{
		cmd:    cmd,
		active: false,
	}
}

func (c *Component) GetCommand() provider.Command {
	return c.cmd
}

func (c *Component) Activate() tea.Cmd {
	c.active = true
	logger.Debug("Activated command", "command", c.cmd.Name)
	return nil
}

func (c *Component) Deactivate() tea.Cmd {
	c.active = false
	logger.Debug("Deactivated command", "command", c.cmd.Name)
	return nil
}

func (c *Component) Resize(width, height int) {
	c.width = width
	c.height = height
}

func (c *Component) Render() string {
	line := theme.BaseStyle().Width(c.width)
	sn := theme.BaseStyle().PaddingRight(2)
	sd := theme.BaseStyle().Foreground(theme.Muted())

	name := c.cmd.Name 
	description := c.cmd.Description

	if c.active {
		name = ansi.Strip(name)
		description = ansi.Strip(description)
		line = line.Foreground(theme.Selection()).Background(theme.Selection())
		sn = sn.Foreground(theme.Foreground()).Background(theme.Selection())
		sd = sd.Foreground(theme.Foreground()).Background(theme.Selection())
	}

	content := lipgloss.JoinHorizontal(
		lipgloss.Top,
		sn.Render(name),
		sd.Render(description),
	)

	return line.Render(content)
}
