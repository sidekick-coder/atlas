package model

import (
	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
	"github.com/sidekick-coder/atlas/tui/components/container"
	"github.com/sidekick-coder/atlas/tui/features/layer"
	"github.com/sidekick-coder/atlas/tui/features/theme"
)


func (m *model) SetSize(width int, height int) {
	layer.ScreenHeight = height 
	layer.ScreenWidth = width 

	m.width = width
	m.height = height

	m.toolbar.SetWidth(width)
	m.footer.SetWidth(width)

	m.toaster.SetScreenSize(width, height)
}

func (m model) View() tea.View {
	body := container.Create().SetSize(m.width-4, m.screenHeight).SetContent("No screen loaded").SetBorder(theme.Current.Primary).SetMargin(0,2).Render()

	if s, ok := m.screen.GetCurrent(); ok {
		body = s.Render()
	}

	layers := []*lipgloss.Layer{}

	layers = append(layers, lipgloss.NewLayer(m.toolbar.Render()).X(0).Y(0).Z(1))
	layers = append(layers, lipgloss.NewLayer(m.tabbar.Render()).X(0).Y(3).Z(1))

	layers = append(layers, lipgloss.NewLayer(body).X(0).Y(4).Z(0))

	layers = append(layers, lipgloss.NewLayer(m.footer.Render()).X(0).Y(m.height - 3).Z(1))

	layers = append(layers, layer.GetLipglossLayers()...)

	if m.toaster.Active {
		layers = append(layers, m.toaster.RenderLayer())
	}

	output := lipgloss.NewCompositor(layers...).Render()

	v := tea.NewView(output)
	v.AltScreen = true

	return v
}
