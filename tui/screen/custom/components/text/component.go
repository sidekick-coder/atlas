package text

import (
	"log"

	"github.com/sidekick-coder/atlas/tui/components/viewport"
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
		width:   p.Width,
		height:  p.Height,
		options: p.Options,
		viewport: viewport,
	}
}

func (c *Component) Render() string {
	return c.viewport.Render()
}

func (c *Component) OnFocus() {}

func (c *Component) OnBlur() {}
