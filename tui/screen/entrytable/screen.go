package entrytable

import (
	tea "charm.land/bubbletea/v2"
	"github.com/sidekick-coder/atlas/internal/app"
	"github.com/sidekick-coder/atlas/tui/app/screen"
	"github.com/sidekick-coder/atlas/tui/components/container"
	"github.com/sidekick-coder/atlas/tui/components/inputdialog"
	"github.com/sidekick-coder/atlas/tui/components/table"
	"github.com/sidekick-coder/atlas/tui/components/toast"
	"github.com/sidekick-coder/atlas/tui/features/chain"
	"github.com/sidekick-coder/atlas/tui/features/context"
	"github.com/sidekick-coder/atlas/tui/features/entryloader"
	"github.com/sidekick-coder/atlas/tui/features/selection"
	tuimodels "github.com/sidekick-coder/atlas/tui/models"
)

type Screen struct {
	app     *app.App
	options map[string]any
	width   int
	height  int

	openScreen  string
	openOptions map[string]any

	selection *selection.Feature
	ctx       *context.Feature
	loader    *entryloader.Feature

	table     *table.Component
	container *container.Component
	dialog    *inputdialog.Component
}

func Create(p tuimodels.ScreenPayload) (tuimodels.Screen, error) {
	openScreen := "entry_single"

	if os, ok := p.Options["open_screen"].(string); ok {
		openScreen = os
	}

	s := &Screen{
		app:     p.App,
		options: p.Options,
		width:   100,
		height:  100,

		openScreen: openScreen,

		loader:    entryloader.Create(*p.App.EntryRepo()),
		selection: selection.Create(),
		ctx:       context.Create(),

		table:     table.Create(),
		container: container.Create(),
		dialog:    inputdialog.Create(),
	}

	return s, nil
}

func (s *Screen) Title() string {
	pt, ok := s.options["title"].(string)

	if ok {
		return pt
	}

	return "entries"
}

func (s *Screen) OpenEntry(cursor int) tea.Cmd {
	entry, ok := s.loader.GetEntry(cursor)

	if !ok {
		return nil
	}

	return screen.Add(s.openScreen, map[string]any{
		"entry": entry.ToMap(),
	})
}

func (s *Screen) OnSubmit(value string) tea.Cmd {
	s.dialog.Close()
	s.loader.SetQuery([]string{value})
	err := s.loader.Load()

	if err != nil {
		return toast.Error(err.Error())
	}

	return nil
}

func (s *Screen) InitDialog() tea.Cmd {
	s.dialog.SetTitle("Search")
	s.dialog.OnSubmit(s.OnSubmit)
	return s.dialog.Init()
}

func (s *Screen) Init() tea.Cmd {
	s.table.OnSelect(s.OpenEntry)
	s.table.SetItemSelection(s.selection)

	q := s.options["query"]

	if q != nil {
		s.loader.SetQuery([]string{q.(string)})
	}

	return chain.Init(
		s.loader.Init,
		s.LoadBindings,
		s.LoadColumns,
		s.table.Init,
		s.InitDialog,
		s.InitSelection,
		s.ctx.Init,
		s.ctx.Activate,
	)
}

func (s *Screen) Dispose() tea.Cmd {
	return chain.Dispose(
		s.DisposeSelection,
		s.table.Dispose,
		s.dialog.Dispose,
		s.UnloadBindings,
		s.ctx.Dispose,
	)
}
