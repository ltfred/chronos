package cmd

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/ltfred/chronos/internal"
	"github.com/ltfred/chronos/models"
	"github.com/spf13/cobra"
)

// configCmd represents the config command
var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Edit config file",
	Run:   config,
}

func init() {
	rootCmd.AddCommand(configCmd)
}

func config(cmd *cobra.Command, args []string) {
	model := models.NewConfigModel(internal.OriginCfg)
	t := tea.NewProgram(model, tea.WithAltScreen())
	_, err := t.Run()
	cobra.CheckErr(err)
}
