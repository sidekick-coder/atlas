package entrysingle

import (
	"fmt"
	"path/filepath"

	tea "charm.land/bubbletea/v2"
	"github.com/sidekick-coder/atlas/internal/app"
	"github.com/sidekick-coder/atlas/internal/models"
	"github.com/sidekick-coder/atlas/internal/utils/maputil"
	"github.com/sidekick-coder/atlas/tui/components/entrymeta"
	"github.com/sidekick-coder/atlas/tui/features/chain"
	"github.com/sidekick-coder/atlas/tui/messages"
	tuimodels "github.com/sidekick-coder/atlas/tui/models"
)

type Screen struct {
	App    *app.App
	Width  int
	Height int
	Path   string
	meta   *entrymeta.Component
	Entry  *models.Entry
	Metas  map[string]string
}

func Create(p tuimodels.ScreenPayload) (tuimodels.Screen, error) {
	path := ""

	if p, ok := p.Options["path"].(string); ok {
		path = p
	}

	if p, ok := maputil.GetString(p.Options, "entry.path"); ok {
		path = p
	}

	if path == "" {
		return nil, fmt.Errorf("path option is required for entrysingle screen")
	}

	e, err := p.App.EntryRepo().GetByPath(path)

	if err != nil {
		return nil, fmt.Errorf("failed to load entry by path: %w", err)
	}

	emc := entrymeta.Create(p.App, path)

	s := &Screen{
		App:    p.App,
		Path:   path,
		meta:   emc,
		Entry:  e,
		Width:  100,
		Height: 100,
		Metas:  map[string]string{},
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

	return chain.Init(s.meta.Init, s.meta.Init)
}

func (s *Screen) Dispose() tea.Cmd {
	return chain.Dispose(s.UnloadBindings, s.meta.Dispose)
}
