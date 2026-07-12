package custom

import (
	"charm.land/lipgloss/v2"
	"github.com/sidekick-coder/atlas/tui/features/theme"
)

func (s *Screen) SetSize(width, height int) {
	s.width = width
	s.height = height

	s.container.
		SetSize(width-4, height).
		SetMargin(0, 2, 0, 2)
}

func (s *Screen) Render() string {

	layers := []*lipgloss.Layer{}

	for _, c := range s.components {
		card := s.card.
			SetSize(c.Cols, c.Rows).
			SetBorder(theme.Current.Secondary).
			SetContent(c.Definition.Render()).
			Render()

		layer := lipgloss.NewLayer(card).
			X(c.X).
			Y(c.Y)

		layers = append(layers, layer)
	}

	content := lipgloss.NewCompositor(layers...).Render()

	s.container.SetContent(content)

	return s.container.Render()
}
