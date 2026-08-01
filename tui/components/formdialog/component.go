package formdialog

import (
	tea "charm.land/bubbletea/v2"
	"github.com/sidekick-coder/atlas/internal/config"
	"github.com/sidekick-coder/atlas/tui/components/dialog"
	"github.com/sidekick-coder/atlas/tui/components/form"
	"github.com/sidekick-coder/atlas/tui/features/chain"
)

type Component struct {
	dialog *dialog.Component
	form   *form.Component
}

func Create() *Component {
	c := &Component{}

	c.form = form.Create()

	c.dialog = dialog.Create()
	c.Resize(c.dialog.GetSize())
	c.dialog.SetContentRender(c.form.Render)

	c.dialog.Events.Open.On(func(){
		c.form.Activate()
	})

	c.dialog.Events.Close.On(func(){
		c.form.Deactivate()
	})

	c.form.Events.Cancel.On(func(){
		c.dialog.Close()
	})

	c.form.Events.Submit.On(func(){
		c.dialog.Close()
	})

	return c
}

func (c *Component) Init() tea.Cmd {
	return chain.Init(c.dialog.Init)
}

func (c *Component) Update(msg tea.Msg) tea.Cmd {
	return chain.Update(
		msg,
		chain.OnKey(c.HadleBinding),
		c.form.Update,
		c.dialog.Update,
	)
}

func (c *Component) Dispose() tea.Cmd {
	return chain.Dispose(c.dialog.Dispose)
}

func (c *Component) Open() {
	c.dialog.Open()
}

func (c *Component) Close() {
	c.dialog.Close()
}

func (c *Component) IsOpen() bool {
	return c.dialog.IsOpen()
}


func (c *Component) SetAction(action config.Action) {
	c.form.SetAction(action)
}
