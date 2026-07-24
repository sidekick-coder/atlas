package entrysingle

import (
	tea "charm.land/bubbletea/v2"
	"github.com/sidekick-coder/atlas/tui/app/screen"
)

func (s *Screen) SetSize(width, height int) {
	s.Width = width
	s.Height = height

	s.EntryMetaComponent.SetSize(s.Width, s.Height)
}

func (s *Screen) HandleSize(msg tea.Msg) tea.Cmd {
	if ss, ok := msg.(screen.SizeMsg); ok {
		s.SetSize(ss.Width, ss.Height)
	}

	return nil
}

func (s *Screen) Render() string {
	return s.EntryMetaComponent.Render()
}
