package cfg

import (
	"fmt"
	"os"
	"runtime"

	"github.com/BurntSushi/toml"
	"github.com/flevin58/fin/tools"
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
type FinToml struct {
	Os      string         `toml:"-"`
	Global  Global         `toml:"global"`
	Darwin  OpSysSpecific  `toml:"darwin"`
	Linux   OpSysSpecific  `toml:"linux"`
	Windows OpSysSpecific  `toml:"windows"`
	Local   *OpSysSpecific `toml:"-"`
}

var (
	Fin FinToml
)

func init() {
	Fin.LoadCfg()
	Fin.Os = runtime.GOOS

	// Here the localized stuff
	switch Fin.Os {
	case "darwin":
		Fin.Local = &Fin.Darwin
	case "linux":
		Fin.Local = &Fin.Linux
	case "windows":
		Fin.Local = &Fin.Windows
	default:
		fmt.Println("I'm sorry but your operating system is not supported yet")
		os.Exit(1)
	}
}

func (f *FinToml) GetTotalApps() []string {
	list := f.Global.Apps
	return append(list, f.Local.Apps...)
}

func (f *FinToml) AddLocalApps(apps ...string) {
	f.Local.Apps = append(f.Local.Apps, apps...)
	f.Local.Apps = tools.Unique(f.Local.Apps)
}

func (f *FinToml) AddGlobalApps(apps ...string) {
	f.Global.Apps = append(f.Global.Apps, apps...)
	f.Global.Apps = tools.Unique(f.Global.Apps)
}

func (f *FinToml) RemoveLocalApps(apps ...string) {
	for _, app := range apps {
		i, found := tools.FindIndexOf(f.Local.Apps, app)
		if found {
			f.Local.Apps = tools.RemoveAtIndex(f.Local.Apps, i)
		}
	}
}

func (f *FinToml) RemoveGlobalApps(apps ...string) {
	for _, app := range apps {
		i, found := tools.FindIndexOf(f.Global.Apps, app)
		if found {
			f.Local.Apps = tools.RemoveAtIndex(f.Global.Apps, i)
		}
	}
}

func (f *FinToml) RemoveApps(apps ...string) {
	f.RemoveGlobalApps(apps...)
	f.RemoveLocalApps(apps...)
}

func (f *FinToml) LoadCfg() {
	data, err := os.ReadFile("fin.toml")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	err = toml.Unmarshal(data, f)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}

func (f *FinToml) SaveCfg() {
	data, err := toml.Marshal(f)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	os.WriteFile("fin.toml", data, 0666)
}
