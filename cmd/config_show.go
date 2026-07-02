/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"charm.land/lipgloss/v2"
	"github.com/spf13/cobra"
	"github.com/sidekick-coder/atlas/internal/config"
)

// config:showCmd represents the config:show command
var configShowCmd = &cobra.Command{
	Use:   "config:show",
	Short: "Prints the current configuration settings",
	Run: func(cmd *cobra.Command, args []string) {
		config := config.Load()

		for key, value := range config.Entries {
			s := lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("12"))

			fmt.Printf("%s: %s\n", s.Render(key), value)
		}
	},
}

func init() {
	rootCmd.AddCommand(configShowCmd)
}
