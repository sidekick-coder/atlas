package inputdialog

import (
	tea "charm.land/bubbletea/v2"
	"github.com/sidekick-coder/atlas/tui/components/dialog"
	"github.com/sidekick-coder/atlas/tui/components/input"
	"github.com/sidekick-coder/atlas/tui/features/chain"
)

type Component struct {
	dialog *dialog.Component
	input *input.Component
	onSubmit func(string) tea.Cmd
}

func Create() *Component {
	return &Component{
		input: input.Create(),
		dialog: dialog.Create(),
	}
}

func (c *Component) OnSubmit(f func(string) tea.Cmd) *Component {
	c.onSubmit = f
	return c
}

func (c *Component) SetTitle(title string) *Component {
	c.dialog.SetTitle(title)
	return c
}

func (c *Component) SetSize(width, height int) *Component {
	c.dialog.SetSize(width, height)
	return c
}


func (c *Component) Open() {
	c.dialog.Open()
}

func (c *Component) SetContent(content string) {
	c.input.SetValue(content)
}

func (c *Component) Close() {
	c.dialog.Close()
}

func (c *Component) IsOpen() bool {
	return c.dialog.IsOpen()
}

func (c *Component) submit() tea.Cmd {
	value := c.input.GetValue()

	if c.onSubmit != nil {
		return c.onSubmit(value)
	}

	c.dialog.Close()

	return nil
}

func (c *Component) Init() tea.Cmd {
	c.dialog.OnRender(c.input.Render)

	c.dialog.OnOpen(func() {
		c.input.Enable()
		c.LoadBindings()
	})

	c.dialog.OnClose(func() {
		c.input.Disable()
		c.UnloadBindings()
	})

	return chain.Init(c.dialog.Init)
}

func (c *Component) Update(msg tea.Msg) tea.Cmd {
	return chain.Update(
		msg,
		chain.OnKey(c.HadleBinding),
		c.input.Update,
		c.dialog.Update,
	)
}

func (c *Component) Dispose() tea.Cmd {
	return chain.Dispose(c.dialog.Dispose)
}
