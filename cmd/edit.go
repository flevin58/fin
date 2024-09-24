/*
Copyright © 2024 Fernando Julio Levin <flevin58@gmail.com>
*/
package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"

	"github.com/spf13/cobra"
)

var editor string

// editCmd represents the edit command
var editCmd = &cobra.Command{
	Use:   "edit",
	Args:  cobra.NoArgs,
	Short: "Edits the fin.toml file using your default editor",
	Long: `Edits the fin.toml file using your default editor
defined in the $EDITOR environment variable.
If none is found, then default OS editors will be chosen:

- Windows: notepad.exe
- Macos: nano
- Linux: nano

To change it either modify the fin.toml file or set the $EDITOR variable
`,
	Run: func(cmd *cobra.Command, args []string) {

		// To debug Fin contents
		// fmt.Printf("Fin: %+v\n", cfg.Fin)
		// os.Exit(0)

		// First determine wich editor to use
		if editor == "" {
			switch runtime.GOOS {
			case "darwin":
				editor = "nano"
			case "windows":
				editor = "notepad.exe"
			case "linux":
				editor = "nano"
			default:
				fmt.Println("I'm sorry but your operating system is not supported yet")
				os.Exit(1)
			}
		}

		// Now get the absolute path to the editor
		editorPath, err := exec.LookPath(editor)
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}
		fmt.Println("Found editor at:", editorPath)

		// Now we can edit the configuration file
		edit := exec.Command(editorPath, "fin.toml")
		edit.Stdin = os.Stdin
		edit.Stdout = os.Stdout
		edit.Stderr = os.Stderr
		err = edit.Run()
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(editCmd)
	editCmd.Flags().StringVarP(&editor, "editor", "e", os.Getenv("EDITOR"), "Use the editor of your choice")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// editCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// editCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
