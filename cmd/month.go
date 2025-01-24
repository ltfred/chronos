package cmd

import (
	"fmt"
	"os"
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
	monthModel := models.NewMonthModel(time.Now().Year())
	p := tea.NewProgram(monthModel)
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}
