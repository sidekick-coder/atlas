package cmd

import (
	"fmt"
	"slices"

	"github.com/spf13/cobra"
	"charm.land/lipgloss/v2"
	"github.com/sidekick-coder/atlas/internal/config"
	"github.com/sidekick-coder/atlas/internal/metadata"
	"github.com/sidekick-coder/atlas/internal/drive/v2"
)

var entryExtractCmd = &cobra.Command{
	Use:   "extract",
	Short: "Extract an entry from the workspace",
	Run: func(cmd *cobra.Command, args []string) {
		config, err := config.Create()

		if err != nil {
			fmt.Println("Error creating config:", err)
			return
		}

		drive, err := drive.New(config.Get("workspace.path"))

		if err != nil {
			fmt.Println("Error creating drive:", err)
			return
		}

		entry, err := drive.Get(args[0])

		if err != nil {
			fmt.Println("Error getting entry:", err)
			return
		}

		handlers := metadata.GetHandlers(entry)


		data, err := metadata.Extract(entry, handlers)

		if err != nil {
			fmt.Println("Error extracting metadata:", err)
			return
		}

		s := lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("12"))

		fmt.Printf("%s\n", entry.Path)

		keys := make([]string, 0, len(data))

		for key := range data {
			keys = append(keys, key)
		}

		slices.SortFunc(keys, func(a, b string) int {
			return len(a) - len(b)
		})

		for _, key := range keys {
			fmt.Printf("%s%s\n", s.Render(key + ":"), data[key])
		}
	},
}

func init() {
	rootCmd.AddCommand(entryExtractCmd)
}
