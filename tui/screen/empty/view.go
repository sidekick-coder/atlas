package empty

import "github.com/sidekick-coder/atlas/tui/features/theme"

func (s *Screen) SetSize(width, height int) {
	s.Width = width
	s.Height = height

	s.list.SetSize(width, 1)

	s.container.
		SetSize(width-4, height).
		SetBorder(theme.Current.Primary).
		SetMargin(0, 2, 0, 2)
}

func (s *Screen) Render() string {
	list := s.list.Render()

	s.container.SetContent(list)

	return s.container.Render()
}
