package root

import (
	"fmt"
	"log"

	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
	"github.com/sidekick-coder/atlas/tui/screen/empty"
)


func (m *model) SetSize(width int, height int) {
	m.width = width
	m.height = height

	m.tabBar.SetWidth(width)
	m.toolbar.SetWidth(width)
	m.footer.SetWidth(width)

	m.input.SetScreenSize(width, height)
	m.toaster.SetScreenSize(width, height)

	sh := height - m.toolbar.GetHeight() - m.tabBar.GetHeight() - m.footer.GetHeight()

	m.screenContainer.SetSize(width, sh)
	m.screenHeight = sh

	log.Printf("SetSize: width=%d, height=%d, screenHeight=%d", width, height, sh)

	for _, s := range m.screens {
		s.SetSize(width, sh)
	}
}

func (m model) View() tea.View {
	body := empty.Placeholder(empty.PlaceholderPayload{
		Width:  m.width,
		Height: m.screenContainer.GetHeight(),
		Text: fmt.Sprintf("No screens available. Press [%s] to add a new screen.", ScreenBindings.Add.Help().Key),
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

	layers := []*lipgloss.Layer{
		lipgloss.NewLayer(content),
	}

	if (m.input.Active) {
		layers = append(layers, m.input.RenderLayer())
	}

	if m.toaster.Active {
		layers = append(layers, m.toaster.RenderLayer())
	}

	output := lipgloss.NewCompositor(layers...).Render()

	v := tea.NewView(output)
	v.AltScreen = true

	return v
}
