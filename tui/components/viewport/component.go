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
func clamp(v, minV, maxV int) int {
	if v < minV {
		return minV
	}
	if v > maxV {
		return maxV
	}
	return v
}

func (c *Component) Render() string {
	lines := strings.Split(c.content, "\n")

	c.offsetY = clamp(c.offsetY, 0, max(0, len(lines)-c.height))
	c.offsetX = clamp(c.offsetX, 0, max(0, c.countX-c.width))

	out := make([]string, 0, c.height)

	for y := c.offsetY; y < min(c.offsetY+c.height, len(lines)); y++ {
		r := []rune(lines[y])

		start := min(c.offsetX, len(r))
		end := min(start+c.width, len(r))

		out = append(out, string(r[start:end]))
	}

	return strings.Join(out, "\n")
}
