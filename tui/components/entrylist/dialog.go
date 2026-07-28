package entrylist

import (
	tea "charm.land/bubbletea/v2"
	"github.com/sidekick-coder/atlas/tui/components/toast"
)

func (c *Component) OnDialogSubmit(value string) tea.Cmd {
	c.dialog.Close()

	c.loader.SetQuery([]string{value})

	err := c.loader.Load()

	if err != nil {
		return toast.Error(err.Error())
	}

	return nil
}

func (c *Component) InitDialog() tea.Cmd {
	c.dialog.SetTitle("Search")
	c.dialog.OnSubmit(c.OnDialogSubmit)

	return c.dialog.Init()
}

