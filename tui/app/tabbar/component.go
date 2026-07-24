package tabbar

import (
	"fmt"

	lipgloss "charm.land/lipgloss/v2"
	"github.com/sidekick-coder/atlas/tui/app/screen"
	"github.com/sidekick-coder/atlas/tui/components/viewport"
	"github.com/sidekick-coder/atlas/tui/features/theme"
)

// Component renders a horizontal tab strip.
type Component struct {
	screen   *screen.Feature
	viewport *viewport.Component
	width    int
}

func Create() *Component {
	return &Component{
		viewport: viewport.Create(),
	}
}

func (t *Component) SetScreen(s *screen.Feature) {
	t.screen = s
}

func (c *Component) SetWidth(width int) {
	c.width = width
}

func (t *Component) Render() string {
	container := theme.BaseStyle().
		Width(t.width).
		Height(1)

	items := make([]string, len(t.screen.GetScreens()))
	screens := t.screen.GetScreens()

	shared := theme.BaseStyle().Padding(0, 1)

	normal := shared.Background(lipgloss.Color(theme.Current.Muted)).Foreground(lipgloss.Color(theme.Current.Background))

	if len(screens) == 0 {
		return container.Render(normal.Render("No tabs"))
	}

	visible := min(len(screens), 5)
	barWidth := max(1, t.width-6)
	tabWidth := barWidth / visible

	indexStyle := theme.BaseStyle().Background(lipgloss.Color(theme.Current.Muted)).Foreground(lipgloss.Color(theme.Current.Background))
	indexActive := theme.BaseStyle().Background(lipgloss.Color(theme.Current.Active)).Foreground(lipgloss.Color(theme.Current.Background))

	textStyle := theme.BaseStyle().
		Background(lipgloss.Color(theme.Current.Placeholder)).
		Foreground(lipgloss.Color(theme.Current.Foreground)).
		MaxWidth(tabWidth-4).
		Padding(0, 1)

	for i, item := range screens {
		title := item.Title()
		id := fmt.Sprintf(" %d ", i)

		row := indexStyle.Render(id) + textStyle.Render(title)

		if i == t.screen.GetCurrentIndex() {
			row = indexActive.Render(id) + textStyle.Render(title)
		}

		items[i] = row
	}


	content := lipgloss.JoinHorizontal(lipgloss.Top, items...)

	return container.Render(content)
}

func (t *Component) GetHeight() int {
	return lipgloss.Height(t.Render())
}

func (t *Component) GetWidth() int {
	return lipgloss.Width(t.Render())
}
