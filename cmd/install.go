/*
Copyright © 2024 Fernando Julio Levin <flevin58@gmail.com>
*/
package cmd

import (
	"fmt"

	"github.com/flevin58/fin/tools/installer"
	"github.com/spf13/cobra"
)

// installCmd represents the install command
var installCmd = &cobra.Command{
	Use:   "install",
	Short: "Installs the given app(s)",
	Long: `Installs the given app(s) passed as arguments
	To install all the configured apps pass all as the argument`,
	Run: func(cmd *cobra.Command, args []string) {
		for _, app := range args {
			fmt.Printf("Installing %s using %s\n", app, installer.Name)
			installer.Install(app)
		}
	},
}

func init() {
	rootCmd.AddCommand(installCmd)
}
