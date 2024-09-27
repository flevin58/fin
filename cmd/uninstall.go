/*
Copyright © 2024 Fernando Julio Levin <flevin58@gmail.com>
*/
package cmd

import (
	"fmt"

	"github.com/flevin58/fin/cfg"
	"github.com/flevin58/fin/tools/installer"
	"github.com/spf13/cobra"
)

// uninstallCmd represents the uninstall command
var uninstallCmd = &cobra.Command{
	Use:   "uninstall",
	Args:  cobra.MinimumNArgs(1),
	Short: "Uninstalls the given app(s)",
	Long:  `Uninstalls the given app(s)`,
	Run: func(cmd *cobra.Command, args []string) {

		// Uninstall given apps
		for _, app := range args {
			fmt.Printf("Uninstalling %s using %s\n", app, installer.Name)
			err := installer.Uninstall(app)
			if flagRemove && err == nil {
				fmt.Printf("Removing %v\n", app)
				cfg.RemoveApps(app)
			}
		}
		cfg.SaveCfg()
	},
}

func init() {
	rootCmd.AddCommand(uninstallCmd)
	uninstallCmd.Flags().BoolVarP(&flagRemove, "remove", "r", false, "The app will be also removed from fin.toml")
}
