package empty

import (
	tea "charm.land/bubbletea/v2"
	"github.com/sidekick-coder/atlas/tui/app/screen"
	"github.com/sidekick-coder/atlas/tui/features/theme"
)

func (s *Screen) SetSize(width, height int) {
	s.Width = width
	s.Height = height

	s.list.SetSize(width, 1)

	s.container.
		SetSize(width-4, height).
		SetBorder(theme.Current.Primary).
		SetMargin(0, 2, 0, 2)
}

func (s *Screen) HandleView(msg tea.Msg) tea.Cmd {
	if msg, ok := msg.(screen.SizeMsg); ok {
		s.SetSize(msg.Width, msg.Height)
	}

	return nil
}

func (s *Screen) Render() string {
	list := s.list.Render()

	s.container.SetContent(list)

	return s.container.Render()
}
