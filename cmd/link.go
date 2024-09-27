/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/flevin58/fin/cfg"
	"github.com/flevin58/fin/tools"
	"github.com/spf13/cobra"
)

// linkCmd represents the link command
var linkCmd = &cobra.Command{
	Use:   "link [flags] source_path dest_path",
	Args:  cobra.MinimumNArgs(1),
	Short: "Creates a symlink of the first arg (source) to the second (target)",
	Long:  `Creates a symlink of the first arg (source) to the second (target)`,
	Run: func(cmd *cobra.Command, args []string) {

		// First handle the case we want to process all links
		if len(args) == 1 && args[0] == "all" {
			for _, link := range cfg.Links {
				src := tools.NormalizePath(link.Src, cfg.Root)
				dst := tools.NormalizePath(link.Dst, "")
				fmt.Printf("Creating symlink %s", dst)

				// If we cannot find source, exit wit error
				// Note that (strangely) os.Symlink does not return error in this case!
				_, err := os.Stat(src)
				if os.IsNotExist(err) {
					fmt.Printf(": can't find %s\n", src)
					return
				}

				err = os.Symlink(src, dst)
				if err != nil {
					fmt.Println(":", err.Error())
				} else {
					fmt.Println(" Ok")
				}
			}
			return
		}

		// Here we process a single link
		src := tools.NormalizePath(args[0], cfg.Root)
		dst := tools.NormalizePath(args[1], "")
		fmt.Printf("Creating symlink %s", dst)
		err := os.Symlink(src, dst)
		if err != nil {
			fmt.Println(":", err.Error())
		} else {
			fmt.Println(" Ok")
		}
	},
}

func init() {
	rootCmd.AddCommand(linkCmd)
	linkCmd.Flags().BoolVarP(&flagAdd, "add", "a", false, "The link will be added to the config list")
}
