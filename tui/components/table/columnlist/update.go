package columnlist

import (
	"strconv"

	tea "charm.land/bubbletea/v2"
	"github.com/sidekick-coder/atlas/tui/components/mapeditor"
	"github.com/sidekick-coder/atlas/tui/features/chain"
)

func (c *Component) InitView() tea.Cmd {
	c.dialog.SetFields([]mapeditor.Field{
		{Label: "Label", FielName: "label"},
		{Label: "Field", FielName: "field"},
		{Label: "Width", FielName: "width"},
	})

	return nil
}

func (c *Component) Update(msg tea.Msg) tea.Cmd {
	return chain.Update(msg, c.dialog.Update, chain.OnKey(c.HandleBindings))
}

func (c *Component) Open() {
	c.LoadBindings()

	c.sidepeeck.Open()
	c.sidepeeck.OnRender(c.Render)

	if c.onOpen != nil {
		c.onOpen()
	}

	c.column.Selection.SetCursor(0)
}

func (c *Component) Close() {
	c.UnloadBindings()

	c.sidepeeck.Close()

	if c.onClose != nil {
		c.onClose()
	}
}

func (c *Component) EditCurrent() tea.Cmd {
	col, ok := c.column.GetColumnSelected()

	if !ok {
		return nil
	}

	width := "auto"

	if col.Width > 0 {
		width = strconv.Itoa(col.Width)
	}

	c.dialog.Open()
	c.dialog.SetValues(map[string]string{
		"label": col.Label,
		"field": col.Field,
		"width": width,
	})

	return nil
}
