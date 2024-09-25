package cfg

import (
	"fmt"
	"os"
	"runtime"

	"github.com/BurntSushi/toml"
)

func loadCfg[T Local | FinToml](tomlfile string) (T, error) {
	var f T
	data, err := os.ReadFile(tomlfile)
	if err != nil {
		return f, err
	}
	err = toml.Unmarshal(data, &f)
	if err != nil {
		return f, err
	}
	return f, nil
}

func saveCfg[T Local | FinToml](cfg *T, tomlfile string) error {
	data, err := toml.Marshal(cfg)
	if err != nil {
		return err
	}
	err = os.WriteFile(tomlfile, data, 0666)
	if err != nil {
		return err
	}
	return nil
}

func LoadCfg() {
	// Load the Local part
	local, err := loadCfg[Local](localTomlPath)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	// Set variables
	Editor = local.Editor
	// See if the editor is already determined
	if Editor == "" {
		switch runtime.GOOS {
		case "darwin":
			Editor = "nano"
		case "windows":
			Editor = "notepad.exe"
		case "linux":
			Editor = "nano"
		default:
			fmt.Printf("I'm sorry but %s is not supported yet", runtime.GOOS)
			os.Exit(1)
		}
	}

	Root = local.Root

	// Load the FinToml structure from the localized file
	fin, err := loadCfg[FinToml](GetTomlPath())
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	Apps = fin.Apps
	Links = fin.Links
}

func SaveCfg() {
	// Save the local data
	local := Local{
		Root:   Root,
		Editor: Editor,
	}
	saveCfg(&local, localTomlPath)

	fin := FinToml{
		Apps:  Apps,
		Links: Links,
	}
	saveCfg(&fin, GetTomlPath())
}
