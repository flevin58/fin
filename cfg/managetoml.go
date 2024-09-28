package cfg

import (
	"os"
	"runtime"

	"github.com/BurntSushi/toml"
	"github.com/flevin58/fin/tools"
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
		tools.ExitWithError(err.Error())
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
			tools.ExitWithError("%s is not yet supported", runtime.GOOS)
		}
	}

	Root = local.Root

	// Load the FinToml structure from the localized file
	fin, err := loadCfg[FinToml](GetTomlPath())
	if err != nil {
		tools.ExitWithError(err.Error())
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
