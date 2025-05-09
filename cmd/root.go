package cmd

import (
	"os"
	"time"

	"github.com/ltfred/chronos/internal"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/ltfred/chronos/models"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "chronos",
	Short: "Chronos is a cli calendar",
	Long:  ``,
	Run:   run,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func run(cmd *cobra.Command, args []string) {
	now := time.Now()
	dayModel := models.NewDayModel(now.Year(), int(now.Month()))
	p := tea.NewProgram(dayModel, tea.WithAltScreen())
	_, err := p.Run()
	cobra.CheckErr(err)
}

func init() {
	cobra.OnInitialize(internal.SetupConfig)
}
