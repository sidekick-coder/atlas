package entrysingle

import (
	"path/filepath"

	tea "charm.land/bubbletea/v2"
	"github.com/sidekick-coder/atlas/internal/app"
)

type Screen struct {
	App *app.App
	Width  int 
	Height int
	Path string
	Metas map[string]string
}

func Create(a *app.App, Path string) *Screen {
	s := &Screen{
		App: a,
		Width:  100,
		Height: 100,
		Path: Path,
		Metas: map[string]string{},
	}

	return s
}


func (s *Screen) Title() string {
	maxLength := 20

	baseName := filepath.Base(s.Path)

	if len(baseName) > maxLength {
		return baseName[:maxLength] + "..."
	}

	return baseName
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
