package cmd

import (
	"fmt"
	"os"

	"github.com/alecthomas/kong"
	"github.com/flevin58/fin/cfg"
	"github.com/flevin58/fin/tools"
)

type CmdLink struct {
	Add bool   `kong:"optional,name='add',help='Adds the symlink to the configuration file'"`
	All bool   `kong:"optional,name='all',help='Process all links in the configuration file'"`
	Src string `kong:"arg,name='source'"`
	Dst string `kong:"arg,name='dest'"`
}

func (c *CmdLink) Run(ctx *kong.Context) error {

	// First handle the case we want to process all links
	if c.All {
		for _, link := range cfg.Links {
			src := tools.NormalizePath(link.Src, cfg.Root)
			dst := tools.NormalizePath(link.Dst, "")
			fmt.Printf("Creating symlink %s ", dst)

			// If we cannot find source, exit wit error
			// Note that (strangely) os.Symlink does not return error in this case!
			_, err := os.Stat(src)
			if err != nil {
				fmt.Println(ErrGliph)
				if os.IsNotExist(err) {
					return fmt.Errorf("can't find %s", src)
				}
			}

			err = os.Symlink(src, dst)
			if err != nil {
				fmt.Println(ErrGliph)
				ctx.Errorf(err.Error())
			} else {
				fmt.Println(OkGliph)
			}
		}
		return nil
	}

	// Here we process a single link
	src := tools.NormalizePath(c.Src, cfg.Root)
	dst := tools.NormalizePath(c.Dst, "")
	fmt.Printf("Creating symlink %s ", dst)
	err := os.Symlink(src, dst)
	if err != nil {
		fmt.Println(ErrGliph)
		ctx.Errorf(err.Error())
	} else {
		fmt.Println(OkGliph)
	}
	return nil
}
