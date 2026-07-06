package entrysingle

import (
	"fmt"
	"path/filepath"

	tea "charm.land/bubbletea/v2"
	"github.com/sidekick-coder/atlas/internal/app"
	"github.com/sidekick-coder/atlas/tui/components/entrymeta"
	"github.com/sidekick-coder/atlas/tui/messages"
	"github.com/sidekick-coder/atlas/tui/models"
)

type Screen struct {
	App                *app.App
	Width              int
	Height             int
	Path               string
	EntryMetaComponent *entrymeta.Component
	Metas              map[string]string
}

func Create(p models.ScreenPayload) (models.Screen, error) {
	path, ok := p.Options["path"].(string)

	if !ok {
		return nil, fmt.Errorf("path option is required for entrysingle screen")
	}

	emc := entrymeta.Create()

	s := &Screen{
		App:                p.App,
		Path:               path,
		EntryMetaComponent: emc,
		Width:              100,
		Height:             100,
		Metas:              map[string]string{},
	}

	return s, nil
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
	err := s.Load()

	if err != nil {
		return messages.ToastErrorCmd(err.Error())
	}


	return nil
}

func (s *Screen) Update(msg tea.Msg) tea.Cmd {
	handlers := []func(tea.Msg) tea.Cmd{}

	handlers = append(handlers, s.HandleScreenKeymaps)

	for _, handler := range handlers {
		cmd := handler(msg)

		if cmd != nil {
			return cmd
		}
	}

	return nil
}


func (s *Screen) SetSize(width, height int) {
	s.Width = width
	s.Height = height

	s.EntryMetaComponent.SetSize(s.Width, s.Height)
}

func (s *Screen) Render() string {
	return s.EntryMetaComponent.
		SetSize(s.Width, s.Height).
		Render()
}
