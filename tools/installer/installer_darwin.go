//go:build darwin

package installer

import (
	"os"
	"os/exec"
)

var (
	installerName     = "brew"
	installerHomePage = "https://brew.sh"
)

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
