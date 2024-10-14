package cmd

import (
	"fmt"
	"runtime"

	"github.com/alecthomas/kong"
	"github.com/flevin58/fin/cfg"
	"github.com/flevin58/fin/tools"
)

type CmdInfo struct {
	Debug bool `kong:"optional,name='debug',short='d',help='Displays some extra stuff'"`
}

func (c *CmdInfo) Run(ctx *kong.Context) error {
	fmt.Println("Fin config file:", cfg.GetTomlPath())
	fmt.Println("Operating System:", runtime.GOOS)
	fmt.Println("Apps List:", cfg.Apps)
	fmt.Println()
	fmt.Println("Installer name:", tools.InstallerName)
	fmt.Println("Installer path:", tools.InstallerPath)
	fmt.Println()
	fmt.Printf("Links: %+v\n", cfg.Links)

	if !c.Debug {
		return nil
	}

	tools.NewTraverse(".").
		WithOnEnterFolder(OnEnter).
		WithOnExitFolder(OnExit).
		WithProcessFile(Process).
		Run()

	return nil
}

func OnEnter(folder string) bool {
	fmt.Println("Entering", folder)
	return folder != ".git"
}

func OnExit(folder string) bool {
	fmt.Println("Leaving", folder)
	return true
}

func Process(file string) bool {
	fmt.Println("Processing", file)
	return true
}
