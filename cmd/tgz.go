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
var tgzCmd = &cobra.Command{
	Use:   "tgz [flags] [source_folder]",
	Args:  cobra.MaximumNArgs(1),
	Short: "Compress a folder to a tgz archive, or list/extract a tgz archive",
	Long: `Compress a folder to a tgz archive, or list/extract a tgz archive.
Examples of usage:
	Compress: fin tgz --output ~/Desktop/myfolder.tgz ./src_folder
	Extract   fin tgz --extract ~/Desktop/myfolder.tgz ./dest_folder
	List:     fin tgz --list ~/Desktop/myfolder.tgz`,
	Run: func(cmd *cobra.Command, args []string) {

		// Case (1): List (no args, --list flag)
		if len(args) == 0 {
			if flagList != "" {
				flagList = tools.NormalizePathWithExt(flagList, "", ".tgz")
				if err := tools.TgzList(flagList); err != nil {
					tools.ExitWithError(err.Error())
				}
				return
			} else {
				tools.ExitWithError("bad command arguments")
			}
			return
		}

		// The arg must be a valid folder
		folder := args[0]
		_, err := os.Stat(folder)
		if err != nil {
			if os.IsNotExist(err) {
				if err = os.MkdirAll(folder, 0644); err != nil {
					tools.ExitWithError(err.Error())
				}
			}
			tools.ExitWithError(err.Error())
		}

		// Case (2): Extract (flag --extract)
		if flagExtract != "" {
			flagExtract = tools.NormalizePath(flagExtract, "")
			if path.Ext(flagExtract) != ".tgz" && !strings.HasSuffix(flagExtract, ".tar.gz") {
				flagExtract += ".tgz"
			}
			fmt.Printf("Extract %s to %s ", flagExtract, folder)
			err := tools.TgzExtract(flagExtract, folder)
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
			flagOutput = tools.NormalizePath(flagOutput, "")
			if path.Ext(flagOutput) != ".tgz" && !strings.HasSuffix(flagOutput, ".tar.gz") {
				flagOutput += ".tgz"
			}
			fmt.Printf("Compressing %s to %s ", folder, flagOutput)
			err := tools.TgzCompress(folder, flagOutput)
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
	rootCmd.AddCommand(tgzCmd)
	tgzCmd.Flags().StringVarP(&flagExtract, "extract", "x", "", "The folder where to extract the given tgz file")
	tgzCmd.Flags().StringVarP(&flagList, "list", "l", "", "List the given tgz file contents")
	tgzCmd.Flags().StringVarP(&flagOutput, "output", "o", "", "The name of the output tgs file")
}
