package custom

import (
	"log/slog"

	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
	"github.com/sidekick-coder/atlas/tui/app/screen"
	"github.com/sidekick-coder/atlas/tui/features/theme"
)

func (s *Screen) HandleSize(msg tea.Msg) tea.Cmd {
	if ss, ok := msg.(screen.SizeMsg); ok {
		slog.Info("custom screen size changed", slog.Int("width", ss.Width), slog.Int("height", ss.Height))
		s.SetSize(ss.Width, ss.Height)
	}

	return nil
}

func (s *Screen) SetSize(width, height int) {
	s.width = width
	s.height = height

	s.container.
		SetSize(width-4, height).
		SetMargin(0, 2, 0, 2)
}

func (s *Screen) Render() string {

	layers := []*lipgloss.Layer{}

	for index, c := range s.components {
		card := s.card.
			SetSize(c.Cols, c.Rows).
			SetBorder(theme.Current.Primary).
			SetContent(c.Definition.Render())

		if s.selection.IsSelected(index) {
			card.SetBorder(theme.Current.Accent)
		}

		layer := lipgloss.NewLayer(card.Render()).
			X(c.X).
			Y(c.Y)

		layers = append(layers, layer)
	}

	content := lipgloss.NewCompositor(layers...).Render()

	s.container.SetContent(content)

	return s.container.Render()
}
