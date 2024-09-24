/*
Copyright © 2024 Fernando Julio Levin <flevin58@gmail.com>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/flevin58/fin/cfg"
	"github.com/flevin58/fin/tools/installer"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "fin",
	Short: "A tool to configure your computer / homepage",
	Long: `A tool to help you manage files installed in your computer,
automatically install / update apps, create symlinks of dotfiles, etc.
All is defined in a ~/.config/fin.toml file
`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Operating System:", cfg.Fin.Os)
		fmt.Println("Global Apps:", cfg.Fin.Global.Apps)
		fmt.Println("Local Apps:", cfg.Fin.Local.Apps)
		fmt.Println()
		fmt.Println("Installer name:", installer.Name)
		fmt.Println("Installer path:", installer.Path)
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
}
