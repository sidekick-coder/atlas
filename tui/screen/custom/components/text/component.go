package text

import (
	tea "charm.land/bubbletea/v2"
	"github.com/sidekick-coder/atlas/tui/components/viewport"
	"github.com/sidekick-coder/atlas/tui/features/chain"
)

type Component struct {
	width   int
	height  int

	viewport *viewport.Component
}

func Create(args ...map[string]any) (*Component, error) {
	props := map[string]any{}

	width := 20
	height := 5

	content := "Text component"

	if len(args) > 0 {
		props = args[0]
	}

	if w, ok := props["width"].(int); ok {
		width = w
	}

	if h, ok := props["height"].(int); ok {
		height = h
	}

	if c, ok := props["content"].(string); ok {
		content = c
	}

	viewport := viewport.Create()

	viewport.SetSize(width, height-2) // 2 padding

	viewport.SetContent(content)

	c := &Component{
		width:    width,
		height:   height,
		viewport: viewport,
	}

	return c, nil
}

func (c *Component) Render() string {
	return c.viewport.Render()
}

func (c *Component) OnFocus() {
	c.LoadBindings()
}

func (c *Component) OnBlur() {
	c.UnloadBindings()
}

func (c *Component) Init() tea.Cmd {
	return nil
}

func (c *Component) Update(msg tea.Msg) tea.Cmd {
	return chain.Update(
		msg,
		chain.OnKey(c.HandleBinding),
	)
}

func (c *Component) Dispose() tea.Cmd {
	return nil
}
