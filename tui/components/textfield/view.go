package textfield

func (c *Component) Resize(width, height int) {
	c.input.SetSize(width-6, 1)
	c.container.SetSize(width-4, 1)
}

func (c *Component) Render() string {
	c.container.SetContent(c.input.Render())

	return c.container.Render()
}
