/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/flevin58/fin/cfg"
	"github.com/flevin58/fin/tools"
	"github.com/flevin58/fin/tools/installer"
	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Args:  cobra.NoArgs,
	Short: "List installed apps",
	Long: `Lists all apps installed with the default installer (brew, apt, scoop).
	You can sync the list with the fin config file with the --sync flag`,
	Run: func(cmd *cobra.Command, args []string) {
		apps, err := installer.List()
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		for i, app := range apps {
			fmt.Printf("%-20s", tools.TrimString(app, 20))
			if i%3 == 0 {
				fmt.Println()
			}
		}
		fmt.Println()

		if flagSync {
			cfg.Apps = apps
			cfg.SaveCfg()
		}
	},
}

var (
	flagSync bool
)

func init() {
	rootCmd.AddCommand(listCmd)
	listCmd.Flags().BoolVarP(&flagSync, "sync", "s", false, "The apps will be synced with the config list")
}
