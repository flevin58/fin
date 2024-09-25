package cfg

import (
	"fmt"
	"os"
	"path"
	"runtime"
	"strings"
)

// Exported vars to other modules
var (
	Root   string
	Editor string
	Apps   []string
	Links  []Link
)

// Internal
var (
	localTomlPath string
)

// This structure is read by GetRootFolder() from  ~/.config/fin.toml
type Local struct {
	Root   string `toml:"root"`   // The root folder for all the installation data
	Editor string `toml:"editor"` // The preferred editor to use
}

// Structure that holds the []links key data
type Link struct {
	Src string `toml:"src"`
	Dst string `toml:"dst"`
}

// Contents of the working fin.toml file (either os localized or determined by -f flag)
type FinToml struct {
	Apps  []string `toml:"apps"`
	Links []Link   `toml:"links"`
}

func GetHomeDir() string {
	homedir, err := os.UserHomeDir()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	return homedir
}

func GetTomlPath() string {
	if strings.HasPrefix(Root, "~/") {
		Root = path.Join("$HOME", Root[2:])
	}
	cfgPath := os.ExpandEnv(Root)
	return path.Join(cfgPath, "fin", fmt.Sprintf("fin_%s.toml", runtime.GOOS))
}

func init() {
	localTomlPath = path.Join(GetHomeDir(), ".config", "fin.toml")
	LoadCfg()
}
