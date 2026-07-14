package dialog

import (
	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
	"github.com/sidekick-coder/atlas/tui/components/borderlabel"
	"github.com/sidekick-coder/atlas/tui/features/chain"
	"github.com/sidekick-coder/atlas/tui/features/layer"
)

type Component struct {
	open   bool
	width  int
	height int
	title  string

	onRender func() string
	onClose  func()
	onOpen   func()

	style  lipgloss.Style
	layer  *layer.Layer
	border *borderlabel.Component
}

func Create() *Component {
	return &Component{
		open:   false,
		width:  100,
		height: 20,
		title:  "",
		border: borderlabel.Create(),

		layer: layer.Create(),
		style: lipgloss.NewStyle().Border(lipgloss.NormalBorder()).Padding(1, 2),
	}
}

func (c *Component) SetTitle(title string) *Component {
	c.title = title
	return c
}

func (c *Component) GetSize() (int, int) {
	return c.width, c.height
}

func (c *Component) SetSize(width, height int) *Component {
	c.width = width
	c.height = height
	return c
}

func (c *Component) SetWidth(width int) *Component {
	c.width = width
	return c
}

func (c *Component) SetPadding(args ...int) *Component {
	c.style = c.style.Padding(args...)
	return c
}

func (c *Component) OnRender(f func() string) *Component {
	c.onRender = f
	return c
}

func (c *Component) Open() {
	c.open = true
	c.LoadBindings()

	if c.onOpen != nil {
		c.onOpen()
	}
}

func (c *Component) Close() {
	c.open = false
	c.UnloadBindings()

	if c.onClose != nil {
		c.onClose()
	}
}

func (c *Component) OnClose(f func()) *Component {
	c.onClose = f
	return c
}

func (c *Component) OnOpen(f func()) *Component {
	c.onOpen = f
	return c
}

func (c *Component) IsOpen() bool {
	return c.open
}

func (c *Component) Init() tea.Cmd {
	x := (layer.ScreenWidth - c.width) / 2
	y := (layer.ScreenHeight - c.height) / 2

	c.layer.SetPosition(x, y)
	c.layer.SetRender(c.render)

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
