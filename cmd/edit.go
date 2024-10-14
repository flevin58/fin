package cmd

import (
	"os"
	"os/exec"

	"github.com/alecthomas/kong"
	"github.com/flevin58/fin/cfg"
)

type CmdEdit struct {
	Editor string `kong:"optional,name='editor',short='e',help='Use the editor of your choice'"`
}

func (c *CmdEdit) Run(ctx *kong.Context) error {

	// If flag -e was given, change the default editor
	if c.Editor != "" {
		// Check if editor exists by getting its absolute path
		if _, err := exec.LookPath(c.Editor); err != nil {
			return err
		}
		// Found it, so update config file
		cfg.Editor = c.Editor
		cfg.SaveCfg()
	}

	// We get the editor from the config file
	editorPath, err := exec.LookPath(cfg.Editor)
	if err != nil {
		return err
	}

	// Now we can edit the configuration file
	edit := exec.Command(editorPath, cfg.GetTomlPath())
	edit.Stdin = os.Stdin
	edit.Stdout = os.Stdout
	edit.Stderr = os.Stderr
	err = edit.Run()
	return err
}
