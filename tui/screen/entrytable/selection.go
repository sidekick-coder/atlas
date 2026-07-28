package entrytable

import (
	tea "charm.land/bubbletea/v2"
	"github.com/sidekick-coder/atlas/tui/features/keymaps"
	"github.com/sidekick-coder/atlas/tui/features/selection"
)

func (s *Screen) OnSelectionChange(event selection.ChangeEvent) {
	s.LoadTrigger()
}

func (s *Screen) InitSelection() tea.Cmd {
	s.selection.SetCursor(-1)
	s.selection.Change.On(s.OnSelectionChange)
	return nil
}

func (s *Screen) LoadTrigger() {
	keymaps.RemoveTriggerByContextID(s.ctx.GetID())

	entry, ok := s.loader.GetEntry(s.selection.GetCursor())

	if ok {
		s.ctx.Set("entry", entry.ToMap())
		trigger := keymaps.CreateEntryTrigger(entry)
		trigger.ContextID = s.ctx.GetID()
		keymaps.AddTrigger(trigger)
	}
}

func (s *Screen) DisposeSelection() tea.Cmd {
	keymaps.RemoveTriggerByContextID(s.ctx.GetID())
	return nil
}
