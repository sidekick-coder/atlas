package entrymeta

func (c *Component) SetSize(width, height int) *Component {
	c.keyValue.SetSize(width, height)
	return c
}

func (c *Component) Render() string {
	mm := map[string]string{}

	for _, m := range c.metas {
		mm[m.Name] = m.Value
	}

	return c.keyValue.Render()
}
