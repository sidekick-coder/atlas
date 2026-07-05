package root

import (
	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
	"github.com/sidekick-coder/atlas/tui/components"
	"github.com/sidekick-coder/atlas/tui/screen/empty"
)


func (m model) View() tea.View {

	body := empty.Placeholder(empty.PlaceholderPayload{
		Width:  m.width,
		Height: m.screenHeight,
		Text:  "No screens available. Press 'a' to add a new screen.",
	})

	if s, ok := m.GetCurrentScreen(); ok {
		body = s.Render()
	}

	content := lipgloss.JoinVertical(
		lipgloss.Left,
		m.toolbar.Render(),
		m.tabBar.Render(),
		body,
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
