/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/sidekick-coder/atlas/internal/config"
	"github.com/spf13/cobra"
)

// config:showCmd represents the config:show command
var configKeymaps = &cobra.Command{
	Use:   "config:keymaps",
	Short: "Prints the current configuration settings",
	Run: func(cmd *cobra.Command, args []string) {
		config, err := config.Create()

		if err != nil {
			fmt.Println("Error creating config:", err)
			return
		}

		entries := config.GetKeymaps()

		for id, k := range entries {
			fmt.Printf("- ID: %s\n", strconv.Itoa(id))
			fmt.Printf("- Keys: %s\n", strings.Join(k.Keys, ", "))
			fmt.Printf("- Description: %s\n", k.Description)
			fmt.Printf("- Keys: %s\n", strings.Join(k.Keys, ", "))

			for _, g := range k.Groups {
				fmt.Printf("- Group: %v\n", g)
			}

			fmt.Printf("- Action: %v\n", k.Action)

			if len(k.ActionOptions) > 0 {
				fmt.Printf("- ActionOptions: %v\n", k.ActionOptions)
			}

			fmt.Printf("\n")

		}
	},
}

func init() {
	rootCmd.AddCommand(configKeymaps)
}
