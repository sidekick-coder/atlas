package toolbar

import (
	"slices"

	tea "charm.land/bubbletea/v2"
	lipgloss "charm.land/lipgloss/v2"
	"github.com/sidekick-coder/atlas/internal/app"
	"github.com/sidekick-coder/atlas/tui/features/key"
	"github.com/sidekick-coder/atlas/tui/features/theme"
)

type Component struct {
	app    *app.App
	path   string
	width  int
	height int
}

func Create(a *app.App) *Component {
	path, _ := a.Config().Get("workspace.path")

	return &Component{
		app:    a,
		path:   path,
		width:  100,
		height: 3,
	}
}

func (c *Component) Update(msg tea.Msg) tea.Cmd {
	if wm, ok := msg.(tea.WindowSizeMsg); ok {
		c.width = wm.Width
	}

	return nil
}

func (c *Component) Render() string {
	iconStyle := lipgloss.NewStyle().Foreground(lipgloss.Color(theme.Current.Primary))

	pending := key.GetPendingTokens()

	if slices.Contains(pending, "<leader>") {
		iconStyle = iconStyle.Foreground(lipgloss.Color(theme.Current.Warning))
	}

	icon := iconStyle.Render("")

	pathStyle := lipgloss.NewStyle().
		BorderStyle(lipgloss.NormalBorder()).
		Width(c.width-4).
		Margin(0, 2, 0, 2).
		Padding(0, 2).
		Height(1).
		BorderForeground(lipgloss.Color(theme.Current.Primary))

	return pathStyle.Render(icon + "  " + c.path)
}
