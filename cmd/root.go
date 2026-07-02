package cmd

import (
	"os"
	"github.com/sidekick-coder/atlas/internal/workspace"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "atlas",
	Short: "Metadata for your files and directories",
	Long: `This is a CLI tool that allows you to add metadata to your files and directories. and then search for them based on that metadata. It is a simple and easy to use tool that can be used to organize your files and directories.`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	  rootCmd.PersistentFlags().StringVarP(
        &workspace.Path,
        "workspace",
		"w",
        "",
        "Workspace directory (defaults to current directory)",
    )

	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}


