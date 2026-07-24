package entrymeta

import (
	"fmt"
	"slices"
	"strings"

	"github.com/sidekick-coder/atlas/internal/models"
)

func (c *Component) Load() error {
	repo := c.app.EntryMetaRepo()

	metas, err := repo.ListByEntryPath(c.path)

	if err != nil {
		return fmt.Errorf("failed to load metadata for entry %s: %w", c.path, err)
	}

	// sort
	slices.SortFunc(metas, func(a, b models.EntryMeta) int {
		if len(a.Name) != len(b.Name) {
			return len(a.Name) - len(b.Name)
		}

		return strings.Compare(a.Name, b.Name)
	})

	c.SetMetas(metas)

	return nil
}
