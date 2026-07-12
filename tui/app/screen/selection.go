package screen

import (
	"fmt"
	"log/slog"
	"github.com/sidekick-coder/atlas/tui/models"
)

func (f *Feature) SetCurrent(index int) error {
	s, ok := f.GetScreenByIndex(index)

	if !ok {
		return fmt.Errorf("invalid screen index: %d", index)
	}

	os, ok := f.GetCurrent()

	if ok {
		slog.Info("close current screen", slog.Int("index", f.Selection.GetCursor()), slog.String("title", os.Title()))
		os.Dispose()
	}

	f.Selection.SetCursor(index)

	s.Init()

	slog.Info("set current screen", slog.Int("index", index), slog.String("title", s.Title()))

	return nil
}

func (f *Feature) GetScreenByIndex(index int) (models.Screen, bool) {
	if index < 0 || index >= len(f.screens) {
		return nil, false
	}

	return f.screens[index], true
}

func (f *Feature) Next() {
	f.SetCurrent(f.Selection.GetNextIndex())
}

func (f *Feature) Prev() {
	f.SetCurrent(f.Selection.GetPrevIndex())
}
