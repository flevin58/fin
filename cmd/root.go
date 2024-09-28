/*
Copyright © 2024 Fernando Julio Levin <flevin58@gmail.com>
*/
package cmd

import (
	"fmt"
	"runtime"

	"github.com/flevin58/fin/cfg"
	"github.com/flevin58/fin/tools"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "fin",
	Args:  cobra.NoArgs,
	Short: "A tool to configure your computer / homepage",
	Long: `A tool to help you manage files installed in your computer,
automatically install / update apps, create symlinks of dotfiles, etc.
All is defined in a ~/.config/fin.toml file
`,
	Run: func(cmd *cobra.Command, args []string) {
		if flagDebug {
			fmt.Println("Fin config file:", cfg.GetTomlPath())
			fmt.Println("Operating System:", runtime.GOOS)
			fmt.Println("Apps List:", cfg.Apps)
			fmt.Println()
			fmt.Println("Installer name:", tools.InstallerName)
			fmt.Println("Installer path:", tools.InstallerPath)
			fmt.Println()
			fmt.Printf("Links: %+v\n", cfg.Links)

			tools.NewTraverse(".").
				WithOnEnterFolder(OnEnter).
				WithOnExitFolder(OnExit).
				WithProcessFile(Process).
				Run()
		}
	},
}

func OnEnter(folder string) bool {
	fmt.Println("Entering", folder)
	return folder != ".git"
}

func OnExit(folder string) bool {
	fmt.Println("Leaving", folder)
	return true
}

func Process(file string) bool {
	fmt.Println("Processing", file)
	return true
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		tools.ExitWithError(err.Error())
	}
}

func init() {
	rootCmd.Flags().BoolVarP(&flagDebug, "debug", "d", false, "We print som more debug info")
}
