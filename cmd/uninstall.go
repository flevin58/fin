package cmd

import (
	"fmt"

	"github.com/alecthomas/kong"
	"github.com/flevin58/fin/cfg"
	"github.com/flevin58/fin/tools"
)

type CmdUninstall struct {
	Remove bool     `kong:"optional,name='remove',short='r',help='The app will be also removed from fin.toml'"`
	Apps   []string `kong:"name='apps',help='Name of apps to be removed'"`
}

// uninstallCmd represents the uninstall command
func (c *CmdUninstall) Run(ctx *kong.Context) error {

	// Uninstall given apps
	for _, app := range c.Apps {
		fmt.Printf("Uninstalling %s using %s\n", app, tools.InstallerName)
		err := tools.Uninstall(app)
		if c.Remove && err == nil {
			fmt.Printf("Removing %v\n", app)
			cfg.RemoveApps(app)
		}
	}
	cfg.SaveCfg()
	return nil
}
