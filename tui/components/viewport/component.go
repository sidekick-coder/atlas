package viewport

import (
	"strings"
)

type Component struct {
	width  int
	height int

	cursorY int
	offsetY int
	maxY  int
	countY int

	content string
}

func Create() *Component {
	return &Component{
		width:   100,
		height:  100,
		offsetY: 0,
		content: "",
	}
}

func (c *Component) SetSize(width, height int) *Component {
	c.width = width
	c.height = height
	return c
}

func (c *Component) SetContent(content string) *Component {
	c.content = content
	c.countY = len(strings.Split(content, "\n"))
	c.maxY = c.countY - c.height
	return c
}

func (c *Component) Up() {
	if c.offsetY > 0 {
		c.offsetY--
	}
}

func (c *Component) Down() {
	if c.offsetY < c.maxY {
		c.offsetY++
	}
}

func (c *Component) Render() string {
	lines := strings.Split(c.content, "\n")
	maxLines := len(lines)

	start := c.offsetY 
	end := min(c.offsetY+c.height, maxLines)

	visible := lines[start:end]

	return strings.Join(visible, "\n")
}
