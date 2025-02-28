package cmd

import (
	"fmt"
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
	p := tea.NewProgram(dayModel)
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(internal.SetupConfig)
}
