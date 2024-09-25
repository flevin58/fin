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

		// First handle the case we want to install all apps
		if len(args) == 1 && args[0] == "all" {
			for _, app := range cfg.Apps {
				fmt.Printf("Installing %s using %s\n", app, installer.Name)
				installer.Install(app)
			}
			return
		}

		// Now we handle the case of having a lis of apps to install
		for _, app := range args {
			fmt.Printf("Installing %s using %s\n", app, installer.Name)
			err := installer.Install(app)
			if err != nil && flagAdd {
				cfg.AddApp(app)
			}
		}
		cfg.SaveCfg()
	},
}

// FLAGS

var (
	flagAdd bool
)

func init() {
	rootCmd.AddCommand(installCmd)
	installCmd.Flags().BoolVarP(&flagAdd, "add", "a", false, "The app will be added to the config list")
}
