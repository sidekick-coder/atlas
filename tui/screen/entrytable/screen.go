package entrytable

import (
	"fmt"
	"log/slog"
	"path/filepath"

	tea "charm.land/bubbletea/v2"
	"github.com/sidekick-coder/atlas/internal/app"
	"github.com/sidekick-coder/atlas/internal/models"
	"github.com/sidekick-coder/atlas/tui/app/program"
	"github.com/sidekick-coder/atlas/tui/app/screen"
	"github.com/sidekick-coder/atlas/tui/components/container"
	"github.com/sidekick-coder/atlas/tui/components/inputdialog"
	"github.com/sidekick-coder/atlas/tui/components/table"
	"github.com/sidekick-coder/atlas/tui/components/toast"
	"github.com/sidekick-coder/atlas/tui/features/chain"
	"github.com/sidekick-coder/atlas/tui/features/entryloader"
	"github.com/sidekick-coder/atlas/tui/features/selection"
	"github.com/sidekick-coder/atlas/tui/features/userkeymaps"
	"github.com/sidekick-coder/atlas/tui/messages"
	tuimodels "github.com/sidekick-coder/atlas/tui/models"
)

type Screen struct {
	app     *app.App
	options map[string]any
	width   int
	height  int

	openScreen  string
	openOptions map[string]any

	selection   *selection.Feature
	loader      *entryloader.Feature
	table       *table.Component
	container   *container.Component
	dialog      *inputdialog.Component
	userKeymaps *userkeymaps.Feature
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

		loader:      entryloader.Create(*p.App.EntryRepo()),
		table:       table.Create(),
		selection:   selection.Create(),
		container:   container.Create(),
		dialog:      inputdialog.Create(),
		userKeymaps: userkeymaps.Create(),
	}

	return s, nil
}

func (s *Screen) Title() string {
	pt, ok := s.options["title"].(string)

	if ok {
		return pt
	}

	return "tables"
}

func (s *Screen) CreateEntryContext(e models.Entry) map[string]any {
	em := e.ToMap()
	em["absolute_path"] = filepath.Join(s.app.Config().GetWorkspaceDir(), e.Path)

	em["update"] = func(payload map[string]any) {
		slog.Info("updating entry", slog.String("path", e.Path), slog.Any("payload", payload))

		for k, v := range payload {
			err := s.app.SetEntryMeta(e.Path, k, fmt.Sprintf("%v", v))

			if err != nil {
				program.Send(messages.ToastErrorMessage(err.Error()))
			}
		}
	}

	return em
}

func (s *Screen) OpenEntry(cursor int) tea.Cmd {

	e, err := s.loader.GetEntry(cursor)

	if err != nil {
		return messages.ToastErrorCmd(err.Error())
	}

	return screen.Add(s.openScreen, map[string]any{
		"entry": s.CreateEntryContext(e),
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

	s.userKeymaps.SetApp(s.app)
	s.userKeymaps.AddGroup("screen_id=entry_table")
	s.userKeymaps.OnBeforeActions(func() error {
		cursor := s.selection.GetCursor()

		if cursor >= 0 {
			e, err := s.loader.GetEntry(cursor)

			if err != nil {
				return err
			}

			s.userKeymaps.AddContext("entry", s.CreateEntryContext(e))
		}

		return nil
	})

	return chain.Init(
		s.loader.Init,
		s.LoadBindings,
		s.LoadColumns,
		s.table.Init,
		s.InitDialog,
	)
}

func (s *Screen) Dispose() tea.Cmd {

	return chain.Dispose(
		s.table.Dispose,
		s.dialog.Dispose,
		s.UnloadBindings,
	)
}
