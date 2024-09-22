/*
Copyright © 2024 Fernando Julio Levin <flevin58@gmail.com>
*/
package cmd

import (
	"fmt"

	"github.com/flevin58/fin/tools/installer"
	"github.com/spf13/cobra"
)

// uninstallCmd represents the uninstall command
var uninstallCmd = &cobra.Command{
	Use:   "uninstall",
	Short: "Uninstalls the given app(s)",
	Long:  `Uninstalls the given app(s)`,
	Run: func(cmd *cobra.Command, args []string) {
		for _, app := range args {
			fmt.Printf("Uninstalling %s using %s\n", app, installer.Name)
			installer.Uninstall(app)
		}
	},
}

func init() {
	rootCmd.AddCommand(uninstallCmd)
}
