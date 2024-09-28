/*
Copyright © 2024 Fernando Julio Levin <flevin58@gmail.com>
*/
package cmd

import (
	"fmt"
	"os"
	"path"
	"strings"

	"github.com/flevin58/fin/tools"

	"github.com/spf13/cobra"
)

// compressCmd represents the compress command
var compressCmd = &cobra.Command{
	Use:   "compress [flags] source_folder",
	Args:  cobra.ExactArgs(1),
	Short: "Compress a folder using zip, tar or tar.gz methods",
	Long: `Compress a folder using zip, tar or tar.gz methods.
The methods are determined by the chosen output flag.
Example:
	fin compress --tar ~/Desktop/myfolder.tar ./myfolder
	fin compress --zip ~/Desktop/myfolder.zip ./myfolder
	fin compress --tgz ~/Desktop/myfolder.tgz ./myfolder`,
	Run: func(cmd *cobra.Command, args []string) {
		folder := args[0]

		// The arg must be a valid folder
		if stat, err := os.Stat(folder); err != nil || !stat.IsDir() {
			fmt.Printf("Error, %s is not a folder\n", folder)
			os.Exit(1)
		}

		// Process flag --zip
		if flagZip != "" {
			if path.Ext(flagZip) != ".zip" {
				flagZip = flagZip + ".zip"
			}
			fmt.Printf("Compress %s to %s", folder, flagZip)
			err := tools.ZipSource(folder, flagZip)
			if err != nil {
				fmt.Println(err.Error())
			} else {
				fmt.Println("Ok")
			}
		}

		// Process flag --tgz
		if flagTgz != "" {
			if path.Ext(flagTgz) != ".tgz" && !strings.HasSuffix(flagTgz, ".tar.gz") {
				flagZip = flagZip + ".tgz"
			}
			fmt.Printf("Compress %s to %s", folder, flagTgz)
			err := tools.TgzSource(folder, flagTgz)
			if err != nil {
				fmt.Println(err.Error())
			} else {
				fmt.Println("Ok")
			}
		}

	},
}

func init() {
	rootCmd.AddCommand(compressCmd)
	compressCmd.Flags().StringVarP(&flagZip, "zip", "", "", "Specifies the 'zip' output file")
	compressCmd.Flags().StringVarP(&flagTgz, "tgz", "", "", "Specifies the 'tar gz' output file")
}
