package entry

import (
	tea "charm.land/bubbletea/v2"
	"github.com/charmbracelet/bubbles/key"
)

func (s *Screen) Update(msg tea.Msg) tea.Cmd {
	switch msg := msg.(type) {
	case tea.KeyPressMsg:
		if key.Matches(msg, Bindings.Up) {
			s.List.MoveUp()
			return nil
		}

		if key.Matches(msg, Bindings.Down) {
			s.List.MoveDown()
			return nil
		}
	}

	return nil
}
