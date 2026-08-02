package entrysingle

import (
	"fmt"

	tea "charm.land/bubbletea/v2"
	"github.com/sidekick-coder/atlas/internal/app"
	"github.com/sidekick-coder/atlas/internal/models"
	"github.com/sidekick-coder/atlas/internal/utils/maputil"
	"github.com/sidekick-coder/atlas/tui/components/entrymeta"
	"github.com/sidekick-coder/atlas/tui/features/chain"
	tuimodels "github.com/sidekick-coder/atlas/tui/models"
)

var ID = "entry-single"

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

	s := &Screen{
		App:    p.App,
		Path:   path,
		meta:   entrymeta.Create(),
		Entry:  e,
		Width:  100,
		Height: 100,
		Metas:  map[string]string{},
	}

	return s, nil
}

func (s *Screen) Init() tea.Cmd {
	return chain.Init(
		s.meta.Init,
		s.meta.Activate,
		s.InitMeta,
	)
}

func (s *Screen) Dispose() tea.Cmd {
	return chain.Dispose(
		s.UnloadBindings,
		s.meta.Deactivate,
		s.meta.Dispose,
	)
}

func (s *Screen) InitMeta() tea.Cmd {
	s.meta.SetProps(map[string]any{
		"entry": s.Entry.ToMap(),
	})
	return nil
}
