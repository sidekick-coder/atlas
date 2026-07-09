package entry

import (
	tea "charm.land/bubbletea/v2"
	"github.com/sidekick-coder/atlas/internal/app"
	"github.com/sidekick-coder/atlas/internal/models"
	"github.com/sidekick-coder/atlas/tui/components/entrylist"
	"github.com/sidekick-coder/atlas/tui/messages"
	tuimodels "github.com/sidekick-coder/atlas/tui/models"
)

type Screen struct {
	App *app.App
	Width  int 
	Height int
	Limit int 
	Query string
	Count int
	TotalPages int
	CurrentPage int
	Entries []models.Entry 
	SelectedIndex int
	SelectedEntry models.Entry 
	SelectedEntryMetas map[string]string
	List *entrylist.EntryList
}

func Create(payload tuimodels.ScreenPayload) (tuimodels.Screen, error) {
	list := entrylist.New()

	s := &Screen{
		App: payload.App,
		Width:  100,
		Height: 100,
		Limit:  30,
		Count: 0,
		TotalPages: 0,
		CurrentPage: 1,
		Query:  "",
		List: list,
		Entries: []models.Entry{},
		SelectedIndex: 0,
		SelectedEntryMetas: map[string]string{},
	}

	if (payload.Options["query"] != nil) {
		if query, ok := payload.Options["query"].(string); ok {
			s.Query = query
		}
	}

	return s, nil
}


func (s *Screen) Title() string {
	return "entries"
}

func (s *Screen) Init() tea.Cmd {
	// Initialization logic for the Entry Screen
	err := s.Load()

	if err != nil {
		return messages.ToastErrorCmd(err.Error(), 3 * 1000)
	}

	return nil
}

func (s *Screen) SetSize(width, height int) {
	s.Width = width
	s.Height = height

	s.List.SetSize(width, height)
}


func (s *Screen) Render() string {
	return s.List.Render()
}

func (s *Screen) Update(msg tea.Msg) tea.Cmd {
	handlers := []func(tea.Msg) tea.Cmd{}

	handlers = append(handlers, s.HandleUserKeyMaps, s.HandleScreenKeymaps)

	for _, handler := range handlers {
		cmd := handler(msg)

		if cmd != nil {
			return cmd
		}
	}

	return nil
}
