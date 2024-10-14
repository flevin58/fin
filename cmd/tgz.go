package cmd

import (
	"fmt"
	"os"

	"github.com/alecthomas/kong"
	"github.com/flevin58/fin/tools"
)

// ================
// tgz sub-commands
// ================
type CmdTgz struct {
	Extract  CmdTgzExtract  `kong:"cmd"`
	List     CmdTgzList     `kong:"cmd"`
	Compress CmdTgzCompress `kong:"cmd"`
}

// ================================
// Sub-command tgz list <tgz_file>
// ================================
type CmdTgzList struct {
	TgzFile string `kong:"arg,required,type='existingfile'"`
}

func (c *CmdTgzList) Run(ctx *kong.Context) error {
	if err := tools.TgzList(c.TgzFile); err != nil {
		return err
	}
	return nil
}

// ===========================================
// Sub-command tgz extract <tgz_file> <folder>
// ===========================================
type CmdTgzExtract struct {
	TgzFile string `kong:"arg,required,type='existingfile'"`
	Folder  string `kong:"arg,required,type='path'"`
}

func (c *CmdTgzExtract) Run(ctx *kong.Context) error {
	// Create the target folder if it does not exist
	if !tools.IsValidFolder(c.Folder) {
		if err := os.MkdirAll(c.Folder, 0755); err != nil {
			return err
		}
	}

	fmt.Printf("Extract %s to %s ", c.TgzFile, c.Folder)
	err := tools.TgzExtract(c.TgzFile, c.Folder)
	if err != nil {
		fmt.Println(ErrGliph)
		ctx.Fatalf(err.Error())
	} else {
		fmt.Println(OkGliph)
	}
	return nil
}

// ============================================
// Sub-command tgz compress <folder> <file.tgz>
// ============================================
type CmdTgzCompress struct {
	Folder  string `kong:"arg,required,type='existingfolder'"`
	TgzFile string `kong:"arg,required,type='path'"`
}

func (c *CmdTgzCompress) Run(ctx *kong.Context) error {
	fmt.Printf("Compressing %s to %s ", c.Folder, c.TgzFile)
	err := tools.TgzCompress(c.Folder, c.TgzFile)
	if err != nil {
		fmt.Println(ErrGliph)
		return err
	} else {
		fmt.Println(OkGliph)
	}
	return nil
}
