package entrylist

import (
	tea "charm.land/bubbletea/v2"
	"github.com/sidekick-coder/atlas/internal/models"
	"github.com/sidekick-coder/atlas/tui/features/keymaps"
	"github.com/sidekick-coder/atlas/tui/features/selection"
)

func (c *Component) InitSelection() tea.Cmd {
	c.list.SetSelection(c.selection)
	c.selection.SetCursor(-1)

	c.selection.Change.On(func (event selection.ChangeEvent) {
		c.LoadTrigger()
	})

	return nil
}

func (c *Component) GetSelection() *selection.Feature {
	return c.selection
}


func (c *Component) LoadTrigger() {
	keymaps.RemoveTriggerByContextID(c.ctx.GetID())

	current, exists := c.GetCurrent()

	if exists {
		c.ctx.Set("entry", current.ToMap())
		trigger := keymaps.CreateEntryTrigger(current)
		trigger.ContextID = c.ctx.GetID()
		keymaps.AddTrigger(trigger)
	}
}

func (c *Component) GetCurrent() (models.Entry, bool) {
	cursor := c.selection.GetCursor()

	entry, err := c.loader.GetEntry(cursor)

	if err != nil {
		return models.Entry{}, false
	}

	return entry, true
}
