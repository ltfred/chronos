package cmd

import (
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/ltfred/chronos/models"
	"github.com/spf13/cobra"
)

// monthCmd represents the month command
var monthCmd = &cobra.Command{
	Use:   "month",
	Short: "Display the month of the current year",
	Long:  "",
	Run:   runMonth,
}

func init() {
	rootCmd.AddCommand(monthCmd)
}

func runMonth(cmd *cobra.Command, args []string) {
	monthModel := models.NewMonthModel(time.Now().Year(), time.Now().Month())
	p := tea.NewProgram(monthModel, tea.WithAltScreen())
	_, err := p.Run()
	cobra.CheckErr(err)
}
