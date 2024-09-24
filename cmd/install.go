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

// installCmd represents the install command
// When invoked without flags, the apps are installed and added locally
// When invoked with the global flag, they are added to the global list
var installCmd = &cobra.Command{
	Use:   "install",
	Args:  cobra.MinimumNArgs(1),
	Short: "Installs the given app(s)",
	Long: `Installs the given app(s) passed as arguments
To install all the configured apps pass all as the only argument.
To add the file(s) to the fin.toml file specify the -a or -g flag`,
	Run: func(cmd *cobra.Command, args []string) {

		var to_install []string

		if len(args) == 1 && args[0] == "all" {
			// Just "all" was given as parameter, so install both local and global apps
			to_install = cfg.Fin.GetTotalApps()
		} else {
			// Here some apps were given as input.
			// We add them to the list (local or global) and install them
			// Note that we are making sure that there are no duplicates in the list
			to_install = args
			if flagGlobal {
				cfg.Fin.AddGlobalApps(args...)
			} else if flagAdd {
				cfg.Fin.AddLocalApps(args...)
			}
		}

		// Now install them
		for _, app := range to_install {
			fmt.Printf("Installing %s using %s\n", app, installer.Name)
			installer.Install(app)
		}

		cfg.Fin.SaveCfg()
	},
}

// FLAGS

var (
	flagAdd    bool
	flagGlobal bool
)

func init() {
	rootCmd.AddCommand(installCmd)
	installCmd.Flags().BoolVarP(&flagGlobal, "gobal", "g", false, "The app will be added to the global fin.toml list")
	installCmd.Flags().BoolVarP(&flagAdd, "add", "a", false, "The app will be added to the local fin.toml list")
}
