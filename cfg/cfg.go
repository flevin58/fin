package cfg

import (
	"fmt"
	"os"
	"runtime"

	toml "github.com/pelletier/go-toml/v2"
)

// Data common to all Operating Systems
type Global struct {
	Apps []string `toml:"apps"`
}

// Structure shared by all Operating Systems
// Each will have its own specific data
type OpSysSpecific struct {
	Apps []string `toml:"apps"`
}

// It is the actual structure used by other modules!
// Only current Operating System data is available
// This makes the code more readable and common to all OS
type FinLocalized struct {
	Global Global
	Local  OpSysSpecific
}

var (
	Fin FinLocalized
)

func init() {
	data, err := os.ReadFile("fin.toml")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	// This is the actual mapping of the .toml file
	type FinToml struct {
		Global  Global
		Darwin  OpSysSpecific
		Linux   OpSysSpecific
		Windows OpSysSpecific
	}

	var tomlData FinToml
	toml.Unmarshal(data, &tomlData)

	// Here the localized stuff
	Fin.Global = tomlData.Global
	switch runtime.GOOS {
	case "darwin":
		Fin.Local = tomlData.Darwin
	case "linux":
		Fin.Local = tomlData.Linux
	case "windows":
		Fin.Local = tomlData.Windows
	default:
		fmt.Println("I'm sorry but your operating system is not supported yet")
		os.Exit(1)
	}
}
