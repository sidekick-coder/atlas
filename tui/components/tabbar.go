package components

import (
	// "strings"

	lipgloss "charm.land/lipgloss/v2"
)

// TabBar renders a horizontal tab strip.
type TabBar struct {
	Items   []string
	Current int
	Width   int
}

func NewTabBar() *TabBar { return &TabBar{} }

func (t *TabBar) SetItems(items []string) {
	t.Items = items
}

func (t *TabBar) SetCurrent(index int) {
	t.Current = index
}

func (t *TabBar) SetWidth(width int) {
	t.Width = width
}

func (t *TabBar) Add(item string) {
	t.Items = append(t.Items, item)
}

func (t *TabBar) Clear() {
	t.Items = []string{}
	t.Current = 0
}

func (t *TabBar) Render() string {
	container := lipgloss.NewStyle().
		Width(t.Width-4).
		Margin(0, 3)

	items := make([]string, len(t.Items))

	shared := lipgloss.NewStyle().Padding(0, 1)

	normal := shared.Background(lipgloss.Color("0")).Foreground(lipgloss.Color("7"))
	active := shared.Background(lipgloss.Color("12")).Foreground(lipgloss.Color("0"))

	for i, item := range t.Items {
		if i == t.Current {
			items[i] = active.Render(item)
		} else {
			items[i] = normal.Render(item)
		}
	}

	row := lipgloss.JoinHorizontal(lipgloss.Top, items...)

	return container.Render(row)
}

func (t *TabBar) GetHeight() int {
	return lipgloss.Height(t.Render())
}

func (t *TabBar) GetWidth() int {
	return lipgloss.Width(t.Render())
}

