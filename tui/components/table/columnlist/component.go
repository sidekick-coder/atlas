package columnlist

import (
	"log"

	tea "charm.land/bubbletea/v2"
	"github.com/sidekick-coder/atlas/tui/components/mapeditor"
	"github.com/sidekick-coder/atlas/tui/components/sidepeeck"
	"github.com/sidekick-coder/atlas/tui/components/table/column"
	"github.com/sidekick-coder/atlas/tui/features/chain"
)

type Component struct {
	sidepeeck *sidepeeck.Component
	dialog    *mapeditor.Component
	column    *column.Feature
	onOpen    func()
	onClose   func()
}

func Create() *Component {
	return &Component{
		sidepeeck: sidepeeck.Create(),
		dialog:    mapeditor.Create(),
	}
}

func (c *Component) IsOpen() bool {
	return c.sidepeeck.IsOpen()
}

func (c *Component) SetColumn(column *column.Feature) {
	c.column = column
}

func (c *Component) OnOpen(fn func()) {
	c.onOpen = fn
}

func (c *Component) OnClose(fn func()) {
	c.onClose = fn
}

func (c *Component) Init() tea.Cmd {
	c.dialog.OnOpen(func() {
		c.UnloadBindings()
	})

	c.dialog.OnClose(func() {
		c.LoadBindings()
	})

	c.dialog.OnSubmit(func(values map[string]string) {
		log.Println("Submitting values:", values)
		column, err := c.column.ParseMapToColumn(values)

		if err != nil {
			log.Println("Error parsing values to column:", err)
			return
		}

		c.column.UpdateSelectedColumn(column)
	})

	return chain.Init(c.sidepeeck.Init, c.InitView, c.dialog.Init)
}

func (c *Component) Dispose() tea.Cmd {
	return chain.Dispose(c.sidepeeck.Dispose, c.dialog.Dispose)
}
