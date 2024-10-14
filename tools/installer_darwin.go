//go:build darwin

package tools

import (
	"os"
	"os/exec"
)

var (
	installerName     = "brew"
	installerHomePage = "https://brew.sh"
)

func Install(app string) error {
	_, err := launchCmd(false, InstallerPath, "install", app)
	return err
}

func Uninstall(app string) error {
	_, err := launchCmd(false, InstallerPath, "uninstall", app)
	return err
}

func Update(app string) error {
	_, err := launchCmd(false, InstallerPath, "upgrade", app)
	return err
}

func installItself() (string, error) {

	install := exec.Command(
		"/bin/bash",
		"-c",
		"$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/HEAD/install.sh)",
	)
	install.Stdin = os.Stdin
	install.Stdout = os.Stdout
	install.Stderr = os.Stderr
	err := install.Run()
	if err != nil {
		return "", err
	}
	return exec.LookPath(installerName)
}
