package entrysingle

import (
	tea "charm.land/bubbletea/v2"
	"github.com/sidekick-coder/atlas/internal/app"
)

type Screen struct {
	App *app.App
	Width  int 
	Height int
}

func Create(a *app.App, entryID int64) *Screen {
	s := &Screen{
		App: a,
		Width:  100,
		Height: 100,
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
}

func (s *Screen) Load() error {
	return nil
}


func (s *Screen) Render() string {
	return "Hello, World!"
}
