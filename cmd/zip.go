/*
Copyright © 2024 Fernando Julio Levin <flevin58@gmail.com>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/flevin58/fin/tools"
	"github.com/spf13/cobra"
)

// compressCmd represents the compress command
var zipCmd = &cobra.Command{
	Use:   "zip [flags] [source_folder]",
	Args:  cobra.MaximumNArgs(1),
	Short: "Compress a folder to a zip archive, or list/extract a zip archive",
	Long: `Compress a folder to a zip archive, or list/extract a zip archive.
Examples of usage:
	Compress: fin zip --output ~/Desktop/myfolder.zip ./src_folder
	Extract   fin zip --extract ~/Desktop/myfolder.zip ./dest_folder
	List:     fin zip --list ~/Desktop/myfolder.zip`,
	Run: func(cmd *cobra.Command, args []string) {

		// Case (1): List (no args, --list flag)
		if len(args) == 0 {
			if flagList != "" {
				flagList = tools.NormalizePathWithExt(flagList, "", ".zip")
				if err := tools.ZipList(flagList); err != nil {
					tools.ExitWithError(err.Error())
				}
				return
			} else {
				tools.ExitWithError("bad command arguments")
			}
			return
		}

		// Normalize the folder
		folder := tools.NormalizePath(args[0], "")

		// Case (2): Extract (flag --extract)
		if flagExtract != "" {
			// Create the target folder if it does not exist
			if !tools.IsValidFolder(folder) {
				if err := os.MkdirAll(folder, 0644); err != nil {
					tools.ExitWithError(err.Error())
				}
			}

			flagExtract = tools.NormalizePathWithExt(flagExtract, "", ".zip")
			fmt.Printf("Extract %s to %s ", flagExtract, folder)
			err := tools.ZipExtract(flagExtract, folder)
			if err != nil {
				fmt.Println(ErrGliph)
				tools.ExitWithError(err.Error())
			} else {
				fmt.Println(OkGliph)
			}
			return
		}

		// Case (3): Compress (flag --output)
		if flagOutput != "" {
			flagOutput = tools.NormalizePathWithExt(flagOutput, "", ".zip")
			fmt.Printf("Compressing %s to %s ", folder, flagOutput)
			err := tools.ZipCompress(folder, flagOutput)
			if err != nil {
				fmt.Println(ErrGliph)
				tools.ExitWithError(err.Error())
			} else {
				fmt.Println(OkGliph)
			}
			return
		}
	},
}

func init() {
	rootCmd.AddCommand(zipCmd)
	zipCmd.Flags().StringVarP(&flagExtract, "extract", "x", "", "The folder where to extract the given tgz file")
	zipCmd.Flags().StringVarP(&flagList, "list", "l", "", "List the given tgz file contents")
	zipCmd.Flags().StringVarP(&flagOutput, "output", "o", "", "The name of the output tgs file")
}
