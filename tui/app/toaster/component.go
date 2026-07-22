package toaster

import (
	tea "charm.land/bubbletea/v2"
	"github.com/sidekick-coder/atlas/tui/components/toast"
	"github.com/sidekick-coder/atlas/tui/features/layer"
)

type Component struct {
	toast toast.Component
	layer *layer.Layer
}

func Create() *Component {
	return &Component{
		toast: *toast.New(),
		layer: layer.Create(),
	}
}

func (c *Component) LoadPosition() {
	x := (layer.ScreenWidth - c.toast.Width) / 2
	y := (layer.ScreenHeight - c.toast.Height) / 2

	c.layer.SetPosition(x, y)
}

func (c *Component) Init() tea.Cmd {
	c.LoadPosition()
	c.layer.SetID("toaster")
	c.layer.SetRender(c.Render)
	c.layer.SetZIndex(3)
	c.InitAction()

	layer.Add(c.layer)

	return nil
}

func (c *Component) Dispose() tea.Cmd {
	layer.Remove(c.layer)

	return nil
}
