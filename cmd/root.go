/*
Copyright © 2024 Fernando Julio Levin <flevin58@gmail.com>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/flevin58/fin/tools/installer"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "mystuff",
	Short: "A tool to configure your computer / homepage",
	Long: `A tool to help you manage files installed in your computer,
automatically install / update apps, create symlinks of dotfiles, etc.
All is defined in a ~/.config/mystuff.toml file
`,
	Run: func(cmd *cobra.Command, args []string) {
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
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.mystuff.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
