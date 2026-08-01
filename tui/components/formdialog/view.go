package formdialog

import "github.com/sidekick-coder/atlas/tui/components/form"

func (c *Component) SetTitle(title string) *Component {
	c.dialog.SetTitle(title)
	return c
}

func (c *Component) SetFields(fields []form.FieldData) *Component {
	c.form.SetFields(fields)
	return c
}

func (c *Component) Resize(width, height int) {
	c.dialog.SetSize(width, height)
	w, h := c.dialog.GetContentSize()
	c.form.Resize(w, h)
}
