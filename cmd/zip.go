package cmd

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
	"strings"

	"github.com/alecthomas/kong"
	"github.com/flevin58/fin/tools"
)

type CmdZip struct {
	Extract  CmdZipExtract  `kong:"cmd,name='extract',help='Extract the given zip file'"`
	List     CmdZipList     `kong:"cmd,name='list',help='List the given zip file contents'"`
	Compress CmdZipCompress `kong:"cmd,name='compress',help='Compress the given folder to zip file'"`
}

// Sub-command zip list <file.zip>
type CmdZipList struct {
	File string `kong:"arg,required,name='zip_file',type='existingfile'"`
}

func (c *CmdZipList) Run(ctx *kong.Context) error {
	if err := tools.ZipList(c.File); err != nil {
		return err
	}
	return nil
}

// Sub-command zip extract <zip_file> [<folder>]
type CmdZipExtract struct {
	File   string `kong:"arg,required,name='zip_file',type='existingfile'"`
	Folder string `kong:"arg,optional,name='folder',type='path',default='',help='The output folder'"`
}

func (c *CmdZipExtract) Run(ctx *kong.Context) error {
	// Guess the target folder if not given: same path as file but with underscores
	// example: /my/dir/myfile.zip => /my/dir/myfile_zip
	if c.Folder == "" {
		folder_name := strings.ReplaceAll(filepath.Base(c.File), ".", "_")
		c.Folder = path.Join(filepath.Dir(c.File), folder_name)
	}

	// Create the target folder if it does not exist
	if !tools.IsValidFolder(c.Folder) {
		if err := os.MkdirAll(c.Folder, 0755); err != nil {
			return err
		}
	}

	fmt.Printf("Extract %s to %s ", c.File, c.Folder)
	err := tools.ZipExtract(c.File, c.Folder)
	if err != nil {
		fmt.Println(ErrGliph)
		ctx.Fatalf(err.Error())
	} else {
		fmt.Println(OkGliph)
	}
	return nil
}

// Sub-command zip compress <folder> --output <file.zip>
type CmdZipCompress struct {
	Folder string `kong:"arg,required,name='folder',type='existingfolder',help='The folder to compress'"`
	Output string `kong:"arg,required,name='zip_file',short='o',help='The name of the zip file to create'"`
}

func (c *CmdZipCompress) Run(ctx *kong.Context) error {
	c.Output = tools.NormalizePathWithExt(c.Output, "", ".zip")
	fmt.Printf("Compressing %s to %s ", c.Folder, c.Output)
	err := tools.ZipCompress(c.Folder, c.Output)
	if err != nil {
		fmt.Println(ErrGliph)
		return err
	} else {
		fmt.Println(OkGliph)
	}
	return nil
}
