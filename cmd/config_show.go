/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"maps"
	"slices"
	"strings"

	"charm.land/lipgloss/v2"
	"github.com/sidekick-coder/atlas/internal/config"
	"github.com/spf13/cobra"
)

// config:showCmd represents the config:show command
var configShowCmd = &cobra.Command{
	Use:   "config:show",
	Short: "Prints the current configuration settings",
	Run: func(cmd *cobra.Command, args []string) {
		config, err := config.Create()

		if err != nil {
			fmt.Println("Error creating config:", err)
			return
		}

		entries := config.GetAll()

		keys := slices.Collect(maps.Keys(entries))

		slices.SortFunc(keys, func(a, b string) int {
			if len(a) != len(b) {
				return len(a) - len(b)
			}

			return strings.Compare(a, b)
		})

		for _, key := range keys {
			s := lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("12"))
			v := strings.ReplaceAll(entries[key], "\n", "\\n")

			fmt.Printf("%s: %s\n", s.Render(key), v)
		}
	},
}

func init() {
	rootCmd.AddCommand(configShowCmd)
}
