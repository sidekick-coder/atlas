package command

import (
	tea "charm.land/bubbletea/v2"
	"github.com/sidekick-coder/atlas/internal/utils"
	"github.com/sidekick-coder/atlas/tui/components/borderlabel"
	"github.com/sidekick-coder/atlas/tui/components/command/item"
	"github.com/sidekick-coder/atlas/tui/components/command/provider"
	"github.com/sidekick-coder/atlas/tui/components/command/providers/action"
	"github.com/sidekick-coder/atlas/tui/components/command/providers/entry"
	"github.com/sidekick-coder/atlas/tui/components/command/providers/screen"
	"github.com/sidekick-coder/atlas/tui/components/dialog"
	"github.com/sidekick-coder/atlas/tui/components/textfield"
	"github.com/sidekick-coder/atlas/tui/features/chain"
	"github.com/sidekick-coder/atlas/tui/features/focusmanager"
)

type Component struct {
	dialog    *dialog.Component
	textfield *textfield.Component
	container *borderlabel.Component

	focus *focusmanager.Feature

	providers []provider.Provider

	commands []*item.Component
}

func Create() *Component {
	c := &Component{
		dialog:    dialog.Create(),
		textfield: textfield.Create(),
		container: borderlabel.Create(),

		focus: focusmanager.Create(),

		providers: []provider.Provider{},
		commands:  []*item.Component{},
	}

	c.dialog.SetContentRender(c.Render)
	c.dialog.SetZIndex(5)
	c.dialog.Bindings.Close.RemoveKey("q")
	c.dialog.Bindings.Close.AddKey("<ctrl+q>")

	c.dialog.Events.Open.On(func() {
		c.textfield.Activate()
		c.LoadControlBindings()
	})

	c.dialog.Events.Close.On(func() {
		c.textfield.SetValue("")
		c.textfield.Deactivate()
		c.UnloadControlBindings()
		c.Load()
	})

	debouncer := utils.NewDebouncer(500)

	c.textfield.Events.Change.On(func() {
		if !c.dialog.IsOpen() {
			return
		}

		debouncer.Do(func() {
			c.Load()
		})
	})

	return c
}

func (c *Component) Init() tea.Cmd {
	return chain.Init(c.dialog.Init, c.textfield.Init, c.InitView, c.InitDefaultProviders, c.Load, c.LoadBindings)
}

func (c *Component) Dispose() tea.Cmd {
	return chain.Dispose(c.dialog.Dispose, c.UnloadBindings)
}

func (c *Component) Update(msg tea.Msg) tea.Cmd {
	return chain.Update(msg,
		c.dialog.Update,
		chain.OnKey(c.HadleBinding),
		c.textfield.Update,
	)
}

func (c *Component) InitDefaultProviders() tea.Cmd {
	c.providers = append(
		c.providers,
		screen.Create(),
		action.Create(),
		entry.Create(),
	)
	return nil
}

func (c *Component) Load() tea.Cmd {
	c.focus.Clear()

	commands := []provider.Command{}

	payload := provider.ListPayload{
		Query: c.textfield.GetValue(),
	}

	for _, p := range c.providers {
		commands = append(commands, p.List(payload)...)
	}

	items := []*item.Component{}

	w, _ := c.dialog.GetContentSize()

	width := w - 1

	for _, cmd := range commands {
		item := item.Create(cmd)

		item.Resize(width, 1)

		c.focus.Add(item)

		items = append(items, item)
	}

	c.commands = items

	containerHeight := 4 + len(c.commands)

	c.container.SetHeight(containerHeight)

	c.focus.First()

	return nil
}

func (c *Component) Execute() tea.Cmd {
	index := c.focus.GetIndex()

	if index < 0 || index >= len(c.commands) {
		return nil
	}

	cmd := c.commands[index].GetCommand()

	result := cmd.Execute()

	c.dialog.Close()

	return result
}
