package syncer

import (
	"fmt"
	"strings"
	lipgloss "charm.land/lipgloss/v2"
)

func (s *Screen) RenderEntries() string {
	border := lipgloss.NewStyle().
		BorderStyle(lipgloss.NormalBorder()).
		Width(s.Width-4).
		Height(s.Height-4).
		Margin(0, 2).
		BorderForeground(lipgloss.Color("12"))

	red := lipgloss.NewStyle().
		Width(s.Width - 4).
		Background(lipgloss.Color("12")).
		Foreground(lipgloss.Color("0"))

	green := lipgloss.NewStyle().
		Width(s.Width - 4).
		Foreground(lipgloss.Color("10"))

	items := []string{}

	maxLength := s.Width - 4

	for _, e := range s.Entries {
		path := e.Path

		if len(path) > maxLength {
			path = path[:maxLength] + "..."
		}

		pad := s.Width - 4 - len([]rune(path))

		pad = max(pad, 0)

		row := green.Render(path + strings.Repeat(" ", pad))

		if e.Error != nil {
			row = red.Render(path + strings.Repeat(" ", pad))
		}

		items = append(items, row)
	}

	maxLines := s.Height - 7
	totalLines := len(items)
	visibleItems := items

	if len(items) > maxLines {
		startIndex := totalLines - maxLines
		endIndex := totalLines
		visibleItems = items[startIndex:endIndex]
	}

	pad := maxLines - len(visibleItems)

	for _ = range pad {
		visibleItems = append(visibleItems, strings.Repeat(" ", s.Width-4))
	}

	footer := lipgloss.NewStyle().Foreground(lipgloss.Color("240")).Render(
		fmt.Sprintf("Total: %d | Time %.2f s", len(s.Entries), s.Time.Seconds()),
	)

	visibleItems = append(visibleItems, footer)

	content := lipgloss.JoinVertical(lipgloss.Left, visibleItems...)

	return border.Render(content)

}
