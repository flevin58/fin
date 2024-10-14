/*
Copyright © 2024 Fernando Julio Levin <flevin58@gmail.com>
*/
package cmd

import (
	"fmt"

	"github.com/alecthomas/kong"
	"github.com/flevin58/fin/cfg"
	"github.com/flevin58/fin/tools"
)

type CmdInstall struct {
	All  bool     `kong:"optional,group='one',name='all',help='Install all apps in the configuration file'"`
	Add  bool     `kong:"optional,group='two',name='add',help='Adds the app to the configuration file'"`
	Apps []string `kong:"arg,required,group='two',name='app',help='The app(s) to be installed'"`
}

func (c *CmdInstall) Run(ctx *kong.Context) error {

	// First handle the case we want to install all apps
	if c.All {
		for _, app := range cfg.Apps {
			fmt.Printf("Installing %s ", app)
			err := tools.Install(app)
			if err != nil {
				fmt.Println(ErrGliph)
			} else {
				fmt.Println(OkGliph)
			}
		}
		return nil
	}

	// Now we handle the case of having a list of apps to install
	for _, app := range ctx.Args {
		fmt.Printf("Installing %s ", app)
		err := tools.Install(app)
		if err != nil {
			fmt.Println(ErrGliph)
		} else {
			fmt.Println(OkGliph)
			if c.Add {
				cfg.AddApp(app)
			}
		}
	}
	cfg.SaveCfg()
	return nil
}
