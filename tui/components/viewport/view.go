package viewport

import (
	"strings"
)

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

		// start := min(c.offsetX, len(r))
		// end := min(start+c.width, len(r))

		out = append(out, string(r))
	}

	return strings.Join(out, "\n")
}
