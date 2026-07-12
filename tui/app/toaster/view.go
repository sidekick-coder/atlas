package toaster

func (c *Component) Render() string {
	if (!c.toast.Active) {
		return ""
	}
	return c.toast.Render()
}
