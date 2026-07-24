package entrymeta

import (
	tea "charm.land/bubbletea/v2"
	"github.com/sidekick-coder/atlas/internal/app"
	"github.com/sidekick-coder/atlas/internal/models"
	"github.com/sidekick-coder/atlas/tui/components/inputdialog"
	"github.com/sidekick-coder/atlas/tui/features/chain"
)

type Component struct {
	Width        int
	Height       int
	Focus        bool
	CurrentIndex int
	Metas        []models.EntryMeta

	app    *app.App
	path   string
	dialog *inputdialog.Component
}

func Create(app *app.App, path string) *Component {
	return &Component{
		Width:        100,
		Height:       100,
		CurrentIndex: 0,
		Focus:        false,
		Metas:        []models.EntryMeta{},

		app:    app,
		path:   path,
		dialog: inputdialog.Create(),
	}
}

func (c *Component) SetFocus(focus bool) *Component {
	c.Focus = focus
	return c
}

func (c *Component) SetMetas(metas []models.EntryMeta) {
	c.Metas = metas

	maxIndex := len(metas) - 1

	if c.CurrentIndex > maxIndex {
		c.CurrentIndex = maxIndex
	}
}

func (c *Component) MoveUp() {
	if c.CurrentIndex > 0 {
		c.CurrentIndex--
	}
}

func (c *Component) MoveDown() {
	if c.CurrentIndex < len(c.Metas)-1 {
		c.CurrentIndex++
	}
}

func (c *Component) GetSelected() (models.EntryMeta, bool) {
	if c.CurrentIndex < 0 || c.CurrentIndex >= len(c.Metas) {
		return models.EntryMeta{}, false
	}

	return c.Metas[c.CurrentIndex], true
}

func (c *Component) Init() tea.Cmd {
	return chain.Init(c.LoadBindings, c.dialog.Init)
}

func (c *Component) Dispose() tea.Cmd {
	return chain.Dispose(c.UnloadBindings, c.dialog.Dispose)
}
