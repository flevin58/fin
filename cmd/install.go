package cmd

import (
	"fmt"

	"github.com/alecthomas/kong"
	"github.com/flevin58/fin/cfg"
	"github.com/flevin58/fin/tools"
)

type CmdInstall struct {
	All  bool     `kong:"xor='all',help='Install all the apps in the configuration file'"`
	Add  bool     `kong:"optional,help='Adds the app to the configuration file'"`
	Apps []string `kong:"xor='all',arg,required,name='app',help='The app(s) to be installed'"`
}

func (c *CmdInstall) Help() string {
	return "pippo"
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
