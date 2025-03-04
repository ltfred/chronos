package cmd

import (
	"github.com/ltfred/chronos/internal"
	"github.com/spf13/cobra"
	"os"
	"os/exec"
)

// configCmd represents the config command
var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Edit config file",
	Run:   initConfig,
}

func init() {
	rootCmd.AddCommand(configCmd)
}

func initConfig(cmd *cobra.Command, args []string) {
	command := exec.Command("vim", internal.GetConfigPath())
	command.Stdin = os.Stdin
	command.Stdout = os.Stdout
	command.Stderr = os.Stderr
	cobra.CheckErr(command.Run())
}
