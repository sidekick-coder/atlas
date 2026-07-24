package screen

import (
	"fmt"
	"github.com/sidekick-coder/atlas/internal/app"
	"github.com/sidekick-coder/atlas/tui/features/key"
	"github.com/sidekick-coder/atlas/tui/features/selection"
	"github.com/sidekick-coder/atlas/tui/models"
)

type OpenScreen struct {
	DefinitionID string
	Screen       models.Screen
}

type Feature struct {
	app          *app.App
	windowWidth  int
	windowHeight int
	screens      []OpenScreen
	bindings     []key.Binding // select screens with <leader>[1,2,3...]
	definitions  map[string]models.ScreenFactory

	Selection *selection.Feature
}

func Create() *Feature {
	return &Feature{
		windowWidth:  100,
		windowHeight: 100,
		screens:      []OpenScreen{},
		bindings:     []key.Binding{},
		definitions:  make(map[string]models.ScreenFactory),

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

	os := OpenScreen{
		DefinitionID: name,
		Screen:       s,
	}

	f.screens = append(f.screens, os)

	index := len(f.screens) - 1

	f.SetCurrent(index)

	binding := key.CreateBinding(fmt.Sprintf("<leader>%d", index)).
		SetDescription(fmt.Sprintf("select screen %d", index)).
		SetHelp(fmt.Sprintf("<leader>%d", index)).
		SetHidden(true).
		SetTags("help-dialog=false")

	key.Register(binding)

	f.bindings = append(f.bindings, binding)

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

	if old, ok := f.GetScreenByIndex(index); ok {
		old.Screen.Dispose()
	}

	os := OpenScreen{
		DefinitionID: name,
		Screen:       s,
	}

	f.screens[index] = os

	f.SetCurrent(index)

	return nil, nil
}

func (f *Feature) Remove(index int) error {
	if index < 0 || index >= len(f.screens) {
		return fmt.Errorf("invalid screen index: %d", index)
	}

	f.screens = append(f.screens[:index], f.screens[index+1:]...)

	current := f.Selection.GetCursor()

	if current >= len(f.screens) {
		f.SetCurrent(len(f.screens) - 1)
	}

	binding := f.bindings[index]

	key.Unregister(binding)

	f.bindings = append(f.bindings[:index], f.bindings[index+1:]...)

	return nil
}

func (f *Feature) GetCurrentIndex() int {
	return f.Selection.GetCursor()
}

func (f *Feature) GetCurrent() (OpenScreen, bool) {
	return f.GetScreenByIndex(f.GetCurrentIndex())
}

func (f *Feature) GetScreens() []models.Screen {
	screens := []models.Screen{}

	for _, os := range f.screens {
		screens = append(screens, os.Screen)
	}

	return screens
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
