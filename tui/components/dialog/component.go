package dialog

import (
	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
	"github.com/sidekick-coder/atlas/tui/components/borderlabel"
	"github.com/sidekick-coder/atlas/tui/features/chain"
	"github.com/sidekick-coder/atlas/tui/features/layer"
)

type Component struct {
	open     bool
	noBorder bool
	width    int
	height   int
	title    string

	Events Events

	contentRender func() string

	style  lipgloss.Style
	layer  *layer.Layer
	border *borderlabel.Component
}

func Create() *Component {
	return &Component{
		open:   false,
		noBorder: true,
		width:  100,
		height: 20,
		title:  "",
		contentRender: func() string {
			return "No content"
		},

		Events: *CreateEvents(),

		border: borderlabel.Create(),

		layer: layer.Create(),
		style: lipgloss.NewStyle().Border(lipgloss.NormalBorder()).Padding(1, 2),
	}
}

func (c *Component) Init() tea.Cmd {
	x := (layer.ScreenWidth - c.width) / 2
	y := (layer.ScreenHeight - c.height) / 2

	c.layer.SetPosition(x, y)
	c.layer.SetRender(c.Render)

	c.LoadDefaultStyle()

	layer.Add(c.layer)
	return nil
}

func (c *Component) Update(msg tea.Msg) tea.Cmd {
	return chain.Update(
		msg,
		chain.OnKey(c.HadleBinding),
	)
}
func (c *Component) Dispose() tea.Cmd {
	layer.Remove(c.layer)

	return nil
}

func (c *Component) SetTitle(title string) *Component {
	c.title = title
	return c
}

func (c *Component) SetPadding(args ...int) *Component {
	c.style = c.style.Padding(args...)
	return c
}

func (c *Component) SetContentRender(f func() string) *Component {
	c.contentRender = f
	return c
}

// @Deprecated: Use SetContentRender instead
func (c *Component) OnRender(f func() string) *Component {
	return c.SetContentRender(f)
}

func (c *Component) Open() {
	c.open = true
	c.LoadBindings()

	c.Events.Open.Emit()
}

func (c *Component) Close() {
	c.open = false
	c.UnloadBindings()

	c.Events.Close.Emit()
}

func (c *Component) Toggle() {
	if c.open {
		c.Close()
		return
	}

	c.Open()
}

func (c *Component) IsOpen() bool {
	return c.open
}
