package entry

import (
	tea "charm.land/bubbletea/v2"
	"github.com/sidekick-coder/atlas/internal/app"
	"github.com/sidekick-coder/atlas/internal/repository/entry"
	"github.com/sidekick-coder/atlas/tui/components/entrylist"
)

type Screen struct {
	App *app.App
	Width  int 
	Height int
	Limit int 
	Offset int 
	Query string
	List *entrylist.EntryList
}

func Create(a *app.App) *Screen {
	list := entrylist.New()

	s := &Screen{
		App: a,
		Width:  100,
		Height: 100,
		Limit:  20,
		Offset: 0,
		Query:  "",
		List: list,
	}

	s.Load()

	return s
}


func (s *Screen) Title() string {
	return "Entries"
}

func (s *Screen) Init() tea.Cmd {
	// Initialization logic for the Entry Screen
	return nil
}

func (s *Screen) SetSize(width, height int) {
	s.Width = width
	s.Height = height

	s.List.SetSize(width, height)
}

func (s *Screen) Load() error {
	repo := s.App.EntryRepo()

	options := entry.ListOptions{
		Limit: s.Limit,
		Offset: s.Offset,
	}

	if s.Query != "" {
		options.Query = []string{s.Query}
	}

	entries, err := repo.List(options)

	if err != nil {
		return err
	}

	s.List.SetEntries(entries)

	return nil
}


func (s *Screen) Render() string {
	return s.List.Render()
}
