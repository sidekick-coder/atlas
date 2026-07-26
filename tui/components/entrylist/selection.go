package entrylist

import (
	"github.com/sidekick-coder/atlas/internal/models"
)

func (c *Component) GetCurrent() (models.Entry, bool) {
	cursor := c.selection.GetCursor()

	entry, err := c.loader.GetEntry(cursor)

	if err != nil {
		return models.Entry{}, false
	}

	return entry, true
}
