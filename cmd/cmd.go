package cmd

import (
	"github.com/alecthomas/kong"
)

var cli struct {
	Edit      CmdEdit      `kong:"cmd,help='Edits the fin.toml file using your default editor'"`
	Info      CmdInfo      `kong:"cmd,help='Get information about the system'"`
	Install   CmdInstall   `kong:"cmd,help='Installs the given app(s)'"`
	List      CmdList      `kong:"cmd,help='List installed apps'"`
	Link      CmdLink      `kong:"cmd,help='Creates a symlink of the first arg (source) to the second (target)'"`
	Tgz       CmdTgz       `kong:"cmd,help='Compress a folder to a tgz archive, or list/extract a tgz archive'"`
	Uninstall CmdUninstall `kong:"cmd,help='Uninstalls the given app(s)'"`
	Zip       CmdZip       `kong:"cmd,help='Compress a folder to a zip archive, or list/extract a zip archive'"`
	Test      CmdTest      `kong:"cmd,help='Test argument parsing by Kong'"`
}

func ParseAndRun() {
	ctx := kong.Parse(&cli)
	err := ctx.Run()
	ctx.FatalIfErrorf(err)
}
