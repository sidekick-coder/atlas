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
	c.viewport.SetSize(width-4, 3)
}

func (t *Component) Render() string {
	container := lipgloss.NewStyle().
		Width(t.width-4).
		Margin(0, 3)

	items := make([]string, len(t.screen.GetScreens()))
	screens := t.screen.GetScreens()

	shared := lipgloss.NewStyle().Padding(0, 1)

	normal := shared.Background(lipgloss.Color(theme.Current.Muted)).Foreground(lipgloss.Color(theme.Current.Background))
	active := shared.Background(lipgloss.Color(theme.Current.Primary)).Foreground(lipgloss.Color(theme.Current.Background))

	if len(screens) == 0 {
		return container.Render(normal.Render("No tabs"))
	}

	visible := min(len(screens), 5)
	barWidth := max(1, t.width-6)
	tabWidth := barWidth / visible

	normal = normal.Width(tabWidth)
	active = active.Width(tabWidth)

	maxVisibleTabs := t.width / tabWidth
	currentIndex := t.screen.GetCurrentIndex()

	startIndex := max(0, currentIndex-(maxVisibleTabs/2))
	endIndex := min(len(screens), startIndex+maxVisibleTabs)

	for i, item := range screens {
		if i < startIndex || i >= endIndex {
			continue
		}

		title := item.Title()

		prefix := fmt.Sprintf("[%d] ", i)

		text := prefix + title

		row := normal.Render(text)

		if i == t.screen.GetCurrentIndex() {
			row = active.Render(text)
			items[i] = row 
			continue
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
