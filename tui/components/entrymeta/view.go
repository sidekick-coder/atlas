package entrymeta

// Deprecated: Use Resize instead
func (c *Component) SetSize(width, height int) *Component {
	c.Resize(width, height)
	return c
}

func (c *Component) Resize(width, height int) {
	c.keyValue.SetSize(width, height)
}


func (c *Component) Render() string {
	mm := map[string]string{}

	for _, m := range c.metas {
		mm[m.Name] = m.Value
	}

	return c.keyValue.Render()
}
