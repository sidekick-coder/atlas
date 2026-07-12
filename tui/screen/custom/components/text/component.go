package text

import (
	"log"

	tea "charm.land/bubbletea/v2"
	"github.com/sidekick-coder/atlas/tui/components/viewport"
	"github.com/sidekick-coder/atlas/tui/features/chain"
	"github.com/sidekick-coder/atlas/tui/screen/custom/component"
)

type Component struct {
	width   int
	height  int
	options map[string]any

	viewport *viewport.Component
}

func Create(p component.DefinitionPayload) component.Definition {
	content := "Text component"

	log.Printf("Creating text component with payload: %+v", p)

	if c, ok := p.Options["content"].(string); ok {
		content = c
	}

	viewport := viewport.Create()

	viewport.SetSize(p.Width, p.Height-2) // 2 padding

	viewport.SetContent(content)

	return &Component{
		width:    p.Width,
		height:   p.Height,
		options:  p.Options,
		viewport: viewport,
	}
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
