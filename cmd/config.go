package cmd

import (
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
	homeDir, err := os.UserHomeDir()
	if err != nil {
		os.Exit(1)
	}
	path := homeDir + "/.config/chronos/config.toml"
	command := exec.Command("vim", path)
	command.Stdin = os.Stdin
	command.Stdout = os.Stdout
	command.Stderr = os.Stderr
	if err = command.Run(); err != nil {
		cmd.PrintErrf("Error: %v\n", err)
		os.Exit(1)
	}
}
