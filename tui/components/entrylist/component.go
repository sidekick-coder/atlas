package entrylist

import (
	tea "charm.land/bubbletea/v2"
	"github.com/sidekick-coder/atlas/internal/logger"
	"github.com/sidekick-coder/atlas/tui/app/program"
	"github.com/sidekick-coder/atlas/tui/components/inputdialog"
	"github.com/sidekick-coder/atlas/tui/components/list"
	"github.com/sidekick-coder/atlas/tui/features/chain"
	"github.com/sidekick-coder/atlas/tui/features/context"
	"github.com/sidekick-coder/atlas/tui/features/entryloader"
	"github.com/sidekick-coder/atlas/tui/features/selection"
)

type Component struct {
	props map[string]any

	selection *selection.Feature
	loader    *entryloader.Feature

	list    *list.Component
	dialog  *inputdialog.Component
	ctx     *context.Feature

	Events *Events
}

func Create() *Component {
	app := program.GetApp()
	repo := app.EntryRepo()

	ctx := context.Create()

	ctx.SetLabel("entry_list")

	return &Component{
		props: map[string]any{},

		loader:    entryloader.Create(*repo),
		selection: selection.Create(),

		list:   list.Create(),
		dialog: inputdialog.Create(),
		ctx:    ctx,

		Events: CreateEvents(),
	}
}

func (c *Component) Init() tea.Cmd {
	return chain.Init(
		c.Load,
		c.ctx.Init,
		c.list.Init,
		c.InitDialog,
		c.InitSelection,
	)
}

func (c *Component) Activate() tea.Cmd {
	return chain.Init(
		c.LoadBindings,
		c.ctx.Activate,
		c.list.Activate,
	)
}

func (c *Component) Deactivate() tea.Cmd {
	return chain.Init(
		c.UnloadBindings,
		c.ctx.Deactivate,
		c.list.Deactive,
	)
}

func (c *Component) Dispose() tea.Cmd {
	return chain.Dispose(
		c.list.Dispose,
		c.ctx.Dispose,
		c.dialog.Dispose,
		c.UnloadBindings,
		c.DisposeSelection,
	)
}

func (c *Component) Context() *context.Feature {
	return c.ctx
}

func (c *Component) Load() tea.Cmd {
	logger.Debug("Loading entries...")
	q := c.props["query"]

	if q != nil {
		c.loader.SetQuery([]string{q.(string)})
	}

	c.loader.Load()
	c.selection.SetTotal(len(c.loader.GetEntries()))
	c.LoadTrigger()

	return c.LoadItems()
}


func (c *Component) SetProps(props map[string]any) tea.Cmd {
	c.props = props
	c.ctx.SetAll(props)
	return c.Load()
}
