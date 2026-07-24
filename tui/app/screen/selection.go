package screen

import (
	"fmt"
	"log/slog"

	"github.com/sidekick-coder/atlas/tui/app/program"
	"github.com/sidekick-coder/atlas/tui/features/keymaps"
)

func (f *Feature) SetCurrent(index int) error {
	os, ok := f.GetScreenByIndex(index)

	if !ok {
		return fmt.Errorf("invalid screen index: %d", index)
	}

	old, ok := f.GetCurrent()

	if ok {
		s := old.Screen

		old.Screen.Dispose()

		slog.Info("close current screen", slog.Int("index", f.Selection.GetCursor()), slog.String("title", s.Title()))
	}

	f.Selection.SetCursor(index)

	os.Screen.Init()
	keymaps.AddGroup("screen", []string{"screen=" + os.DefinitionID})

	program.Send(f.Size())

	slog.Info("set current screen", slog.Int("index", index), slog.String("title", os.Screen.Title()))

	return nil
}

func (f *Feature) GetScreenByIndex(index int) (OpenScreen, bool) {
	if index < 0 || index >= len(f.screens) {
		return OpenScreen{}, false
	}

	return f.screens[index], true
}

func (f *Feature) Next() {
	f.SetCurrent(f.Selection.GetNextIndex())
}

func (f *Feature) Prev() {
	f.SetCurrent(f.Selection.GetPrevIndex())
}
