package root

import (
	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
	"github.com/sidekick-coder/atlas/tui/components"
)


func (m model) View() tea.View {
	s, ok := m.GetCurrentScreen()

	if !ok {
		s = m.emptyScreen
	}

	content := lipgloss.JoinVertical(
		lipgloss.Left,
		m.toolbar.Render(),
		m.tabBar.Render(),
		s.Render(),
		m.footer.Render(),
	)

	if components.GlobalToast.Active() {
		content = components.PlaceOverlay(components.GlobalToast.Box(), content, m.width, m.height)
	}

	layers := []*lipgloss.Layer{
		lipgloss.NewLayer(content),
	}

	if (m.input.Active) {
		layers = append(layers, m.input.RenderLayer())
	}

	output := lipgloss.NewCompositor(layers...).Render()

	v := tea.NewView(output)
	v.AltScreen = true

	return v
}
