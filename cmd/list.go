package cmd

import (
	"fmt"

	"github.com/alecthomas/kong"
	"github.com/flevin58/fin/cfg"
	"github.com/flevin58/fin/tools"
)

type CmdList struct {
	Sync bool `kong:"optional,name='sync',short='s',help='Update configuration with detected installed apps'"`
}

func (c *CmdList) Run(ctx *kong.Context) error {
	apps, err := tools.List()
	if err != nil {
		return err
	}

	for i, app := range apps {
		fmt.Printf("%-20s", tools.TrimString(app, 20))
		if i%3 == 0 {
			fmt.Println()
		}
	}
	fmt.Println()

	if c.Sync {
		cfg.Apps = apps
		cfg.SaveCfg()
	}

	return nil
}
