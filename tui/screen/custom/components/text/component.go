package text

import (
	"log"

	"github.com/sidekick-coder/atlas/tui/screen/custom/component"
)

type Component struct {
	width   int
	height  int
	options map[string]any

	content string
}

func Create(p component.DefinitionPayload) component.Definition {
	content := "Text component"

	log.Printf("Creating text component with payload: %+v", p)

	if c, ok := p.Options["content"].(string); ok {
		content = c
	}

	return &Component{
		width:   p.Width,
		height:  p.Height,
		options: p.Options,
		content: content,
	}
}

func (c *Component) Render() string {
	return c.content
}
