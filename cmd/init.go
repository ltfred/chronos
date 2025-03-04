package cmd

import (
	"fmt"
	"github.com/ltfred/chronos/internal"
	"github.com/spf13/cobra"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize the configuration file",
	Run: func(cmd *cobra.Command, args []string) {
		var isOverwrite string
		cmd.Print("The initialization configuration file will overwrite the existing file, continue?[Y/n]")
		_, err := fmt.Scan(&isOverwrite)
		cobra.CheckErr(err)
		if isOverwrite == "Y" || isOverwrite == "y" {
			internal.InitConfigFile()
		} else {
			cmd.Println("Abort")
		}
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}
