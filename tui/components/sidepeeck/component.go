package sidepeeck

import (
	"log/slog"

	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
	"github.com/sidekick-coder/atlas/tui/features/layer"
)

type Component struct {
	open bool 
	width int 
	height int

	onRender func() string

	style lipgloss.Style
	layer *layer.Layer
}

func Create() *Component {
	return &Component{
		open: false,
		width: 100,
		height: 100,

		layer: layer.Create(),
		style: lipgloss.NewStyle().Border(lipgloss.NormalBorder()).Padding(1, 2),
	}
}

func (c *Component) GetSize() (int, int) {
	return c.width, c.height
}

func (c *Component) SetSize(width, height int) {
	c.width = width
	c.height = height
}

func (c *Component) OnRender(f func() string) {
	c.onRender = f
}

func (c *Component) Open() {
	c.open = true
}

func (c *Component) Close() {
	c.open = false
}

func (c *Component) IsOpen() bool {
	return c.open
}

func (c *Component) Init() tea.Cmd {
	c.height = layer.ScreenHeight 

	slog.Info("Initializing sidepeek component", "width", c.width, "height", c.height, "screenWidth", layer.ScreenWidth, "screenHeight", layer.ScreenHeight)

	x := layer.ScreenWidth - c.width

	c.layer.SetPosition(x, 0)
	c.layer.SetZIndex(2)
	c.layer.SetRender(c.Render)
	c.layer.SetID("sidepeek-" + c.layer.ID)

	c.LoadDefaultStyle()

	layer.Add(c.layer)
	return nil
}

func (c *Component) Dispose() tea.Cmd {
	layer.Remove(c.layer)
	return nil
}
