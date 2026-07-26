package text

import (
	tea "charm.land/bubbletea/v2"
	"github.com/sidekick-coder/atlas/tui/components/viewport"
)

type Component struct {
	props map[string]any
	viewport *viewport.Component
}

func Create() *Component {
	return &Component{
		viewport: viewport.Create(),
	}
}

func (c *Component) Init() tea.Cmd {
	return nil
}

func (c *Component) Dispose() tea.Cmd {
	return nil
}

func (c *Component) Activate() tea.Cmd {
	return c.LoadBindings()
}

func (c *Component) Deactivate() tea.Cmd {
	return c.UnloadBindings()
}

func (c *Component) Focus() tea.Cmd {
	return c.Activate()
}

func (c *Component) Blur() tea.Cmd {
	return c.Deactivate()
}

func (c *Component) SetProps(props map[string]any) {
	c.props = props

	if content, ok := props["content"].(string); ok {
		c.viewport.SetContent(content)
	}
}
