package cmd

import (
	"fmt"
	"path"
	"strings"

	"github.com/alecthomas/kong"
	"github.com/flevin58/fin/tools"
)

type CmdTgz struct {
	Extract string `kong:"optional,type='path',name='extract',short='x',help='The folder where to extract the given zip file'"`
	List    string `kong:"optional,existingfile,name='list',short='l',help='List the given zip file contents'"`
	Output  string `kong:"optional,type='path',name='output',short='o',help='The name of the output zip file'"`
	Folder  string `kong:"arg,type='path',help='The name of the folder'"`
}

// compressCmd represents the compress command
func (c *CmdTgz) Run(ctx *kong.Context) error {

	// Case (1): List (no args, --list flag)
	if c.List != "" {
		c.List = tools.NormalizePathWithExt(c.List, "", ".tgz")
		if err := tools.TgzList(c.List); err != nil {
			return err
		}
		return nil
	}

	// Case (2): Extract (flag --extract)
	if c.Extract != "" {
		if path.Ext(c.Extract) != ".tgz" && !strings.HasSuffix(c.Extract, ".tar.gz") {
			c.Extract += ".tgz"
		}
		fmt.Printf("Extract %s to %s ", c.Extract, c.Folder)
		err := tools.TgzExtract(c.Extract, c.Folder)
		if err != nil {
			fmt.Println(ErrGliph)
			return err
		} else {
			fmt.Println(OkGliph)
		}
		return nil
	}

	// Case (3): Compress (flag --output)
	if c.Output != "" {
		if path.Ext(c.Output) != ".tgz" && !strings.HasSuffix(c.Output, ".tar.gz") {
			c.Output += ".tgz"
		}
		fmt.Printf("Compressing %s to %s ", c.Folder, c.Output)
		err := tools.TgzCompress(c.Folder, c.Output)
		if err != nil {
			fmt.Println(ErrGliph)
			return err
		} else {
			fmt.Println(OkGliph)
		}
	}
	return nil
}
