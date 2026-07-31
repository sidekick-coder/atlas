package logs

import (
	tea "charm.land/bubbletea/v2"
	"github.com/sidekick-coder/atlas/tui/features/selection"
)

func (s *Screen) OnSelectionChange(e selection.ChangeEvent){
	s.viewport.ScrollToVisible(e.New)
}

func (s *Screen) InitSelection() tea.Cmd {
	s.selection.Change.On(s.OnSelectionChange)
	return nil
}

