package viewport

import (
	"strings"
)

type Component struct {
	width  int
	height int

	cursorY int
	offsetY int
	maxY    int
	countY  int

	cursorX int
	offsetX int
	maxX    int
	countX  int

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
	lines := strings.Split(content, "\n")

	for _, line := range lines {
		if w := len([]rune(line)); w > c.countX {
			c.countX = w
		}
	}

	c.content = content
	c.countY = len(lines)

	c.maxY = max(0, c.countY-c.height)
	c.maxX = max(0, c.countX-c.width)

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

func (c *Component) Left() {
	if c.offsetX > 0 {
		c.offsetX--
	}
}

func (c *Component) Right() {
	if c.offsetX < c.maxX {
		c.offsetX++
	}
}

func (c *Component) GetOffsetY() int {
	return c.offsetY
}

func (c *Component) SetOffsetY(offsetY int) {
	c.offsetY = clamp(offsetY, 0, c.maxY)
}

func (c *Component) ScrollTo(y int) {
	c.offsetY = clamp(y, 0, c.maxY)
}

func (c *Component) IsVisible(y int) bool {
	return y >= c.offsetY && y < c.offsetY+c.height
}

func (c *Component) ScrollToVisible(y int) {
	if c.IsVisible(y) {
		return
	}

	if y < c.offsetY {
		c.offsetY = y
		return
	}

	if y >= c.offsetY+c.height {
		c.offsetY = y - c.height + 1
	}
}

