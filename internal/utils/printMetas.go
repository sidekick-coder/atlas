package utils

import (
	"fmt"
	"slices"
	"github.com/sidekick-coder/atlas/internal/models"
	"charm.land/lipgloss/v2"
)

func PrintMap(payload map[string]string) {
	s := lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("12"))

	keys := make([]string, 0, len(payload))

	for key := range payload {
		keys = append(keys, key)
	}

	slices.SortFunc(keys, func(a, b string) int {
		return len(a) - len(b)
	})

	for _, key := range keys {
		fmt.Printf("%s: %s\n", s.Render(key), payload[key])
	}


}

func PrintMetas(metas []models.EntryMeta) {
	mapMetas := map[string]string{}

	for _, meta := range metas {
		mapMetas[meta.Name] = meta.Value
	}

	PrintMap(mapMetas)
}
