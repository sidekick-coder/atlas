package entrylist

import (
	tea "charm.land/bubbletea/v2"
	"github.com/sidekick-coder/atlas/tui/app/program"
	"github.com/sidekick-coder/atlas/tui/components/inputdialog"
	"github.com/sidekick-coder/atlas/tui/components/list"
	"github.com/sidekick-coder/atlas/tui/components/toast"
	"github.com/sidekick-coder/atlas/tui/features/chain"
	"github.com/sidekick-coder/atlas/tui/features/entryloader"
	"github.com/sidekick-coder/atlas/tui/features/selection"
)

type Component struct {
	props map[string]any

	selection *selection.Feature
	loader    *entryloader.Feature

	list   *list.Component
	dialog *inputdialog.Component
}

func Create() *Component {
	app := program.GetApp()
	repo := app.EntryRepo()

	return &Component{
		props: map[string]any{},

		loader:    entryloader.Create(*repo),
		selection: selection.Create(),

		list:   list.Create(),
		dialog: inputdialog.Create(),
	}
}

func (c *Component) Init() tea.Cmd {
	return chain.Init(
		c.Load,
		c.list.Init,
		c.InitDialog,
		c.InitSelection,
	)
}

func (c *Component) Activate() tea.Cmd {
	return chain.Init(
		c.LoadBindings,
		c.list.Activate,
	)
}

func (c *Component) Deactivate() tea.Cmd {
	return chain.Init(
		c.UnloadBindings,
		c.list.Deactive,
	)
}

func (c *Component) Dispose() tea.Cmd {
	return chain.Dispose(
		c.list.Dispose,
		c.dialog.Dispose,
		c.UnloadBindings,
	)
}

func (c *Component) OnSubmit(value string) tea.Cmd {
	c.dialog.Close()

	c.loader.SetQuery([]string{value})

	err := c.loader.Load()

	if err != nil {
		return toast.Error(err.Error())
	}

	return nil
}

func (c *Component) InitDialog() tea.Cmd {
	c.dialog.SetTitle("Search")
	c.dialog.OnSubmit(c.OnSubmit)

	return c.dialog.Init()
}

func (c *Component) Load() tea.Cmd {
	q := c.props["query"]

	if q != nil {
		c.loader.SetQuery([]string{q.(string)})
	}

	c.loader.Load()

	return c.LoadItems()
}

func (c *Component) InitSelection() tea.Cmd {
	c.list.SetSelection(c.selection)
	c.selection.SetCursor(-1)
	return nil
}

func (c *Component) Focus() tea.Cmd {
	return c.Activate()
}

func (c *Component) Blur() tea.Cmd {
	return c.Deactivate()
}

func (c *Component) SetProps(props map[string]any) tea.Cmd {
	c.props = props
	return c.Load()
}
