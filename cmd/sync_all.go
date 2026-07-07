package cmd

import (
	"fmt"

	"charm.land/lipgloss/v2"
	"github.com/sidekick-coder/atlas/internal/app"
	"github.com/sidekick-coder/atlas/internal/sync/v2"
	"github.com/spf13/cobra"
)

var syncAllCmd = &cobra.Command{
	Use:   "sync:all",
	Short: "Sync all entries",
	Run: func(cmd *cobra.Command, args []string) {
		concurrency, err := cmd.Flags().GetInt("concurrency")

		app, err := app.Create()

		if err != nil {
			fmt.Println(err)
			return
		}

		syncer := app.Syncer()

		green := lipgloss.NewStyle().Foreground(lipgloss.Color("10"))
		red := lipgloss.NewStyle().Foreground(lipgloss.Color("9"))

		onSuccess := func(path string, metas map[string]string) {
			fmt.Printf("%s\n", green.Render(path))
		}

		onError := func(path string, err error) {
			fmt.Printf("%s\n", red.Render(path))
			fmt.Printf("Error: %v\n", err)
		}

		payload := sync.AllPayload{
			Concurrency: concurrency,
			OnSuccess:   onSuccess,
			OnError:     onError,
		}

		result, err := syncer.All(payload)

		if err != nil {
			fmt.Printf("Error syncing entries: %v\n", err)
			return
		}

		fmt.Printf("\n")
		fmt.Printf("Total Entries: %d\n", result.TotalEntries)
		fmt.Printf("Total Batches: %d\n", result.TotalBatches)
		fmt.Printf("Time: %.2f\n", result.Time.Seconds())
		fmt.Printf("Concurrency: %d\n", result.Concurrency)
	},
}

func init() {
	rootCmd.AddCommand(syncAllCmd)
	syncAllCmd.Flags().IntP("concurrency", "c", 1, "Number of concurrent workers")
}
