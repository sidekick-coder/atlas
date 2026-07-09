package container

import (
	lipgloss "charm.land/lipgloss/v2"
)

type Component struct {
	content string
	width   int
	height  int
	style   lipgloss.Style
}

func Create() *Component {
	c := &Component{
		content: "",
		width:   100,
		height:  100,
		style: lipgloss.NewStyle(),
	}

	return c
}


func (c *Component) SetMargin(i ...int) *Component {
	c.style = c.style.Margin(i...)
	return c
}

func (c *Component) SetPadding(i ...int) *Component {
	c.style = c.style.Padding(i...)
	return c
}

func (c *Component) SetBorder(color string) *Component {
	c.style = c.style.BorderStyle(lipgloss.NormalBorder()).BorderForeground(lipgloss.Color(color))
	return c
}


func (c *Component) SetSize(width, height int) *Component {
	c.style = c.style.Width(width).Height(height)
	c.width = width
	c.height = height
	return c
}

func (c *Component) SetContent(content string) *Component {
	c.content = content
	return c
}

func (c *Component) Render() string {
	return c.style.Render(c.content)
}

func (c *Component) GetHeight() int {
	return lipgloss.Height(c.Render())
}

func (c *Component) GetWidth() int {
	return lipgloss.Width(c.Render())
}
