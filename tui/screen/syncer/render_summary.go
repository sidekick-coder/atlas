package syncer

import (
	"fmt"
	lipgloss "charm.land/lipgloss/v2"
)

func (s *Screen) RenderSummary() string {
	border := lipgloss.NewStyle().
		BorderStyle(lipgloss.NormalBorder()).
		Width(s.Width-4).
		Height(s.Height-4).
		Margin(0, 2).
		BorderForeground(lipgloss.Color("12"))

	warning := lipgloss.NewStyle().
		Foreground(lipgloss.Color("11"))

	green := lipgloss.NewStyle().
		Foreground(lipgloss.Color("10"))

	items := []string{}

	if (s.Running) {
		items = append(items, warning.Render("Syncing..."))
	}

	if (s.Completed) {
		items = append(items, green.Render("Completed!"))
	}

	items = append(items, fmt.Sprintf("Total: %d", s.TotalEntries))
	items = append(items, fmt.Sprintf("Time: %.2f s", s.Time.Seconds()))

	content := lipgloss.JoinVertical(lipgloss.Left, items...)

	content = lipgloss.Place(s.Width-4, s.Height-8, lipgloss.Center, lipgloss.Center, content)

	return border.Render(content)

}
