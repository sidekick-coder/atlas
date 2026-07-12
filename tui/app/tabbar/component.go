package tabbar

import (
	lipgloss "charm.land/lipgloss/v2"
	"github.com/sidekick-coder/atlas/tui/app/screen"
	"github.com/sidekick-coder/atlas/tui/features/theme"
)

// Component renders a horizontal tab strip.
type Component struct {
	screen *screen.Feature
	Width  int
}

func Create() *Component { return &Component{} }

func (t *Component) SetScreen(s *screen.Feature) {
	t.screen = s
}

func (t *Component) SetWidth(width int) {
	t.Width = width
}

func (t *Component) Render() string {
	container := lipgloss.NewStyle().
		Width(t.Width-4).
		Margin(0, 3)

	items := make([]string, len(t.screen.GetScreens()))

	shared := lipgloss.NewStyle().Padding(0, 1)

	normal := shared.Background(lipgloss.Color(theme.Current.Foreground)).Foreground(lipgloss.Color(theme.Current.Background))
	active := shared.Background(lipgloss.Color(theme.Current.Primary)).Foreground(lipgloss.Color(theme.Current.Background))

	for i, item := range t.screen.GetScreens() {
		row := normal.Render(item.Title())

		if i == t.screen.GetCurrentIndex() {
			row = active.Render(item.Title())
		}

		items[i] = row
	}

	if len(items) == 0 {
		return container.Render(normal.Render("No tabs"))
	}

	row := lipgloss.JoinHorizontal(lipgloss.Top, items...)

	return container.Render(row)
}

func (t *Component) GetHeight() int {
	return lipgloss.Height(t.Render())
}

func (t *Component) GetWidth() int {
	return lipgloss.Width(t.Render())
}
