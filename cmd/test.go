package cmd

import (
	"fmt"
	"strings"

	"github.com/alecthomas/kong"
	"github.com/flevin58/fin/cfg"
)

type CmdTest struct {
	All  bool     `kong:"optional,xor='x',help='Install all the apps in the configuration file'"`
	Add  bool     `kong:"optional,xor='x',help='Adds the app to the configuration file'"`
	Apps []string `kong:"arg,optional,name='app',help='The app(s) to be installed'"`
}

func (c *CmdTest) Help() string {
	return "Help string"
}

func (c *CmdTest) Usage() string {
	return "***usage***"
}

func (c *CmdTest) Run(ctx *kong.Context) error {

	// First handle the case we want to install all apps
	if c.All {
		if len(ctx.Args) != 2 {
			return fmt.Errorf("flag --all should be used alone")
		}
		for _, app := range cfg.Apps {
			app = strings.TrimSpace(app)
			if app == "" {
				continue
			}
			fmt.Printf("Installing %s ", app)
			result := true
			if !result {
				fmt.Println(ErrGliph)
			} else {
				fmt.Println(OkGliph)
			}
		}
		return nil
	}

	// Now we handle the case of having a list of apps to install
	for _, app := range c.Apps {
		fmt.Printf("Installing %s ", app)
		result := true
		if !result {
			fmt.Println(ErrGliph)
		} else {
			fmt.Println(OkGliph)
			if c.Add {
				//cfg.AddApp(app)
			}
		}
	}
	cfg.SaveCfg()
	return nil
}
