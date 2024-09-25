/*
Copyright © 2024 Fernando Julio Levin <flevin58@gmail.com>
*/
package cmd

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/flevin58/fin/cfg"
	"github.com/spf13/cobra"
)

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

		// Now get the absolute path to the editor
		editorPath, err := exec.LookPath(editor)
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}

		// If flag -e was given, change the default editor
		if cmd.Flags().Changed("editor") {
			cfg.Editor = editor
			cfg.SaveCfg()
		}

		// Now we can edit the configuration file
		edit := exec.Command(editorPath, cfg.GetTomlPath())
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

// FLAGS

var (
	editor string
)

func init() {
	if cfg.Editor == "" {
		cfg.Editor = os.Getenv("EDITOR")
	}
	rootCmd.AddCommand(editCmd)
	editCmd.Flags().StringVarP(&editor, "editor", "e", cfg.Editor, "Use the editor of your choice")
}
