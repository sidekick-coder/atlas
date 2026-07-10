package borderlabel

import "github.com/sidekick-coder/atlas/tui/features/theme"

type Component struct {
	label string
	content string
	color string
	width int 
	height int
}

func Create() *Component {
	return &Component{
		label: "",
		content: "",
		height: 10,
		width: 10,
		color: theme.Current.Primary,
	}
}

func (c *Component) GetSize() (int, int) {
	return c.width, c.height
}

func (c *Component) SetSize(width, height int) *Component {
	c.width = width
	c.height = height
	return c
}

func (c *Component) GetLabel() string {
	return c.label
}

func (c *Component) SetLabel(label string) *Component {
	c.label = label
	return c
}

func (c *Component) GetContent() string {
	return c.content
}

func (c *Component) SetContent(content string) *Component {
	c.content = content
	return c
}
