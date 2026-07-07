package root

import (
	"fmt"

	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
	"github.com/sidekick-coder/atlas/tui/components"
	"github.com/sidekick-coder/atlas/tui/screen/empty"
)


func (m *model) SetSize(width int, height int) {
	m.width = width
	m.height = height

	components.GlobalInput.SetSize(width, height)
	components.GlobalToast.SetSize(width, height)

	m.tabBar.SetWidth(width)
	m.toolbar.SetWidth(width)
	m.footer.SetWidth(width)

	m.input.SetScreenSize(width, height)
	m.toaster.SetScreenSize(width, height)

	toolbarHeight := 1
	tabBarHeight := 1
	footerHeight := 1
	contentHeight := height - toolbarHeight - tabBarHeight - footerHeight
	m.screenHeight = contentHeight

	for _, s := range m.screens {
		s.SetSize(width, contentHeight)
	}
}

func (m model) View() tea.View {
	m.LoadBindings()

	body := empty.Placeholder(empty.PlaceholderPayload{
		Width:  m.width,
		Height: m.screenHeight,
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

	if components.GlobalToast.Active() {
		content = components.PlaceOverlay(components.GlobalToast.Box(), content, m.width, m.height)
	}

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
