package footer

import (
	"fmt"
	"strings"
	lipgloss "charm.land/lipgloss/v2"
	tkey "github.com/sidekick-coder/atlas/tui/features/key"
	"github.com/sidekick-coder/atlas/tui/features/theme"
)

var (
	footerStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color(theme.Current.Muted))

	footerKeyStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color(theme.Current.Primary)).
			Bold(true)
)

func (f *Component) Render() string {
	container := lipgloss.NewStyle().
		BorderStyle(lipgloss.NormalBorder()).
		Width(f.width-4).
		Margin(0, 2).
		Padding(0, 2).
		BorderForeground(lipgloss.Color(theme.Current.Primary))

	var parts []string

	remaningWidth := f.width

	for _, b := range tkey.GetBindings() {
		k := b.GetHelp()
		d := b.GetDescription()

		if f.dialog.IsOpen() && !b.HasTag("global:help") {
			continue
		}

		if k == "" || d == "" || b.IsHidden() {
			continue
		}

		part := fmt.Sprintf("%s %s", footerKeyStyle.Render(k), footerStyle.Render(d))

		remaningWidth -= lipgloss.Width(part)

		parts = append(parts, fmt.Sprintf("%s %s", footerKeyStyle.Render(k), footerStyle.Render(d)))

		if remaningWidth <= 80 {
			parts = append(parts, footerStyle.Render(fmt.Sprintf("... and %d more", len(tkey.GetBindings())-len(parts))))
			break
		}

	}

	sep := footerStyle.Render(" · ")
	row := strings.Join(parts, sep)

	return container.Render(row)
}

