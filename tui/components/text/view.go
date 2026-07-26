package text

func (c *Component) Render() string {
	return c.viewport.Render()
}

func (c *Component) Resize(width, height int) {
	c.viewport.SetSize(width, height)
}
