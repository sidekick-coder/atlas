package screen

import (
	"fmt"

	"github.com/sidekick-coder/atlas/tui/app/program"
)

type ScreenEntry struct {
	ID      string
	Options map[string]any
}

func List() ([]ScreenEntry, error) {
	config := program.GetConfig()

	entries := []ScreenEntry{}

	entries = append(entries, ScreenEntry{
		ID:      "entry_list",
		Options: map[string]any{},
	})

	entries = append(entries, ScreenEntry{
		ID:      "entry_table",
		Options: map[string]any{},
	})

	entries = append(entries, ScreenEntry{
		ID:      "logs",
		Options: map[string]any{},
	})

	us, err := config.GetScreens()

	if err != nil {
		return nil, fmt.Errorf("error getting user screens: %w", err)
	}

	for _, s := range us {
		entries = append(entries, ScreenEntry{
			ID:      s.ID,
			Options: s.Options,
		})
	}

	return entries, nil
}
