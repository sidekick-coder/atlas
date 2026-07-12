package custom

import (
	"fmt"
	"log/slog"

	tea "charm.land/bubbletea/v2"
	"github.com/sidekick-coder/atlas/internal/template"
	"github.com/sidekick-coder/atlas/tui/components/container"
	"github.com/sidekick-coder/atlas/tui/components/toast"
	"github.com/sidekick-coder/atlas/tui/features/chain"
	"github.com/sidekick-coder/atlas/tui/features/selection"
	"github.com/sidekick-coder/atlas/tui/models"
	"github.com/sidekick-coder/atlas/tui/screen/custom/component"
	"github.com/sidekick-coder/atlas/tui/screen/custom/components/text"
)

type Entry struct {
	ID      string
	Options map[string]any
}

type Screen struct {
	title   string
	width   int
	height  int
	options map[string]any

	components []component.Component
	selection *selection.Feature

	container container.Component
	card      container.Component
}

func Create(p models.ScreenPayload) (models.Screen, error) {
	s := &Screen{
		width:      100,
		height:     100,
		title:      "custom",
		options:    p.Options,

		components: []component.Component{},
		container:  *container.Create(),
		card:       *container.Create(),
		selection: selection.Create(),
	}

	if title, ok := p.Options["title"].(string); ok {
		t, err := template.Render(title, p.Options)

		if err != nil {
			return nil, fmt.Errorf("error rendering title: %w", err)
		}

		s.title = t
	}

	return s, nil
}

func (s *Screen) Title() string {
	return s.title
}

func (s *Screen) LoadComponents() tea.Cmd {
	component.RegisterDefinition("text", text.Create)

	s.components = []component.Component{}

	oc, ok := s.options["components"].([]any)

	if !ok {
		return toast.Error("No components found in options")
	}

	oc, err := template.ParseArray(oc, s.options)

	if err != nil {
		return toast.Error(fmt.Sprintf("Error parsing components: %s", err.Error()))
	}

	for _, c := range oc {
		component, err := component.CreateFromMap(c)

		if err != nil {
			return toast.Error(err.Error())
		}

		s.components = append(s.components, *component)
	}

	s.selection.SetTotal(len(s.components))
	s.Select(0)

	slog.Info("loaded components", slog.Int("count", len(s.components)))

	return nil
}

func (s *Screen) Init() tea.Cmd {
	return chain.Init(s.LoadComponents, s.LoadBindings)
}

func (s *Screen) Dispose() tea.Cmd {
	return chain.Dispose(s.UnloadBindings)
}
