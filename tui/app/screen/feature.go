package screen

import (
	"fmt"
	"log/slog"

	"github.com/sidekick-coder/atlas/internal/app"
	"github.com/sidekick-coder/atlas/tui/features/selection"
	"github.com/sidekick-coder/atlas/tui/models"
)

type Feature struct {
	app         *app.App
	windowWidth  int
	windowHeight int
	screens     []models.Screen
	definitions map[string]models.ScreenFactory

	Selection *selection.Feature
}

func Create() *Feature {
	return &Feature{
		windowWidth:  100,
		windowHeight: 100,
		screens:     []models.Screen{},
		definitions: make(map[string]models.ScreenFactory),

		Selection: selection.Create(),
	}
}

func (f *Feature) SetApp(a *app.App) {
	f.app = a
}

func (f *Feature) CreateScreen(name string, options ...map[string]any) (models.Screen, error) {
	fac, ok := f.definitions[name]

	if !ok {
		return nil, fmt.Errorf("invalid screen name: %s", name)
	}

	payload := models.ScreenPayload{
		App:     f.app,
		Options: make(map[string]any),
	}

	if len(options) > 0 {
		payload.Options = options[0]
	}

	s, err := fac(payload)

	if err != nil {
		return nil, err
	}

	return s, nil
}

func (f *Feature) Add(name string, options ...map[string]any) (models.Screen, error) {
	s, err := f.CreateScreen(name, options...)

	if err != nil {
		return nil, err
	}

	f.screens = append(f.screens, s)

	f.SetCurrent(len(f.screens) - 1)

	slog.Info("added screen", slog.String("name", name))

	f.Selection.SetTotal(len(f.screens))

	return nil, nil
}

func (f *Feature) Replace(index int, name string, options ...map[string]any) (models.Screen, error) {
	if index < 0 || index >= len(f.screens) {
		return nil, fmt.Errorf("invalid screen index: %d", index)
	}

	s, err := f.CreateScreen(name, options...)

	if err != nil {
		return nil, err
	}

	f.screens[index] = s

	f.SetCurrent(index)

	slog.Info("replaced screen", slog.Int("index", index), slog.String("name", name))

	return nil, nil
}

func (f *Feature) Remove(index int) error {
	if index < 0 || index >= len(f.screens) {
		return fmt.Errorf("invalid screen index: %d", index)
	}

	f.screens = append(f.screens[:index], f.screens[index+1:]...)

	current:= f.Selection.GetCursor()

	if current >= len(f.screens) {
		f.SetCurrent(len(f.screens) - 1)
	}

	slog.Info("removed screen", slog.Int("index", index))

	return nil
}

func (f *Feature) GetCurrentIndex() int {
	return f.Selection.GetCursor()
}

func (f *Feature) GetCurrent() (models.Screen, bool) {
	return f.GetScreenByIndex(f.GetCurrentIndex())
}

func (f *Feature) GetScreens() []models.Screen {
	return f.screens
}

func (f *Feature) SetDefinition(id string, fac models.ScreenFactory) {
	f.definitions[id] = fac
}

func (f *Feature) GetDefinition(id string) (models.ScreenFactory, bool) {
	fac, ok := f.definitions[id]

	return fac, ok
}

func (f *Feature) Init() error {
	f.LoadBindings()

	return nil
}
