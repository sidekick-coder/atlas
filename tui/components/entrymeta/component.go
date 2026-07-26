package entrymeta

import (
	"fmt"
	"slices"
	"strings"

	tea "charm.land/bubbletea/v2"
	"github.com/sidekick-coder/atlas/internal/models"
	"github.com/sidekick-coder/atlas/internal/utils/maputil"
	"github.com/sidekick-coder/atlas/tui/app/program"
	"github.com/sidekick-coder/atlas/tui/components/inputdialog"
	"github.com/sidekick-coder/atlas/tui/components/keyvalue"
	"github.com/sidekick-coder/atlas/tui/features/chain"
	"github.com/sidekick-coder/atlas/tui/features/selection"
)

type Component struct {
	metas []models.EntryMeta
	path  string
	props map[string]any

	selection *selection.Feature

	dialog   *inputdialog.Component
	keyValue *keyvalue.Component
}

func Create() *Component {
	return &Component{
		metas: []models.EntryMeta{},
		path:  "",
		props: map[string]any{},

		selection: selection.Create(),

		dialog:   inputdialog.Create(),
		keyValue: keyvalue.Create(),
	}
}

func (c *Component) Activate() tea.Cmd {
	return chain.Cmd(c.LoadBindings, c.keyValue.Init, c.dialog.Init)
}

func (c *Component) Deactivate() tea.Cmd {
	return chain.Cmd(c.UnloadBindings, c.keyValue.Dispose, c.dialog.Dispose)
}

func (c *Component) Focus() tea.Cmd {
	return c.Activate()
}

func (c *Component) Blur() tea.Cmd {
	return c.Deactivate()
}

func (c *Component) Init() tea.Cmd {
	c.keyValue.SetSelection(c.selection)
	return nil
}

func (c *Component) Dispose() tea.Cmd {
	return nil
}

func (c *Component) SetProps(props map[string]any) {
	c.props = props

	if p, ok := props["path"].(string); ok {
		c.path = p
	}

	if p, ok := maputil.GetString(props, "entry.path"); ok {
		c.path = p
	}

	c.Load()
}

func (c *Component) Load() error {
	c.keyValue.Clear()

	if c.path == "" {
		return nil
	}

	app := program.GetApp()

	repo := app.EntryMetaRepo()

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

	c.metas = metas

	items := []keyvalue.Item{}

	for _, meta := range metas {
		items = append(items, keyvalue.Item{
			Key:   meta.Name,
			Value: meta.Value,
		})
	}

	c.keyValue.SetItems(items)

	return nil
}

func (c *Component) GetSelected() (models.EntryMeta, bool) {
	item, ok := c.keyValue.GetSelected()

	if !ok {
		return models.EntryMeta{}, false
	}

	for _, meta := range c.metas {
		if meta.Name == item.Key {
			return meta, true
		}
	}

	return models.EntryMeta{}, false
}
