package cmd

import (
	"fmt"
	"os"

	"github.com/alecthomas/kong"
	"github.com/flevin58/fin/tools"
)

// ================
// zip sub-commands
// ================
type CmdZip struct {
	Extract  CmdZipExtract  `kong:"cmd"`
	List     CmdZipList     `kong:"cmd"`
	Compress CmdZipCompress `kong:"cmd"`
}

// ===============================
// Sub-command zip list <zip-file>
// ===============================
type CmdZipList struct {
	ZipFile string `kong:"arg,required,type='existingfile'"`
}

func (c *CmdZipList) Run(ctx *kong.Context) error {
	if err := tools.ZipList(c.ZipFile); err != nil {
		return err
	}
	return nil
}

// ===========================================
// Sub-command zip extract <zip_file> <folder>
// ===========================================
type CmdZipExtract struct {
	ZipFile string `kong:"arg,required,type='existingfile'"`
	Folder  string `kong:"arg,required,type='path'"`
}

func (c *CmdZipExtract) Run(ctx *kong.Context) error {
	// Create the target folder if it does not exist
	if !tools.IsValidFolder(c.Folder) {
		if err := os.MkdirAll(c.Folder, 0755); err != nil {
			return err
		}
	}

	fmt.Printf("Extract %s to %s ", c.ZipFile, c.Folder)
	err := tools.ZipExtract(c.ZipFile, c.Folder)
	if err != nil {
		fmt.Println(ErrGliph)
		ctx.Fatalf(err.Error())
	} else {
		fmt.Println(OkGliph)
	}
	return nil
}

// ============================================
// Sub-command zip compress <folder> <zip-file>
// ============================================
type CmdZipCompress struct {
	Folder  string `kong:"arg,required,type='existingfolder'"`
	ZipFile string `kong:"arg,required,type='path'"`
}

func (c *CmdZipCompress) Run(ctx *kong.Context) error {
	fmt.Printf("Compressing %s to %s ", c.Folder, c.ZipFile)
	err := tools.ZipCompress(c.Folder, c.ZipFile)
	if err != nil {
		fmt.Println(ErrGliph)
		return err
	} else {
		fmt.Println(OkGliph)
	}
	return nil
}
