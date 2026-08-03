package entrycontroller

import (
	tea "charm.land/bubbletea/v2"
	"github.com/sidekick-coder/atlas/internal/app"
	"github.com/sidekick-coder/atlas/tui/features/chain"
)

type Feature struct {
	app *app.App
}

func Create(app *app.App) *Feature {
	return &Feature{
		app: app,
	}
}

func (f *Feature) Update(msg tea.Msg) tea.Cmd {
	return chain.Update(
		msg, 
		f.HandleSet,
		f.HandleUnset,
		f.HandleSync,
		f.HandleAllSync,
	)
}

func (f *Feature) Init() tea.Cmd {
	return nil
}

func (f *Feature) Dispose() tea.Cmd {
	return nil
}
