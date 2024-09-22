package installer

import (
	"fmt"
	"os"
	"os/exec"
)

var (
	Name = installerName
	Path = GetPath()
)

func launchCmd(cmd string, args ...string) error {
	command := exec.Command(cmd, args...)
	command.Stdin = os.Stdin
	command.Stdout = os.Stdout
	command.Stderr = os.Stderr
	return command.Run()
}

func Install(app string) {
	err := launchCmd(Path, "install", app)
	if err != nil {
		fmt.Printf("Failed to install %s\n", app)
	} else {
		fmt.Printf("App %s successfully installed\n", app)
	}
}

func Uninstall(app string) {
	err := launchCmd(Path, "uninstall", app)
	if err != nil {
		fmt.Printf("Failed to uninstall %s\n", app)
	} else {
		fmt.Printf("App %s successfully uninstalled\n", app)
	}
}

func Update(app string) {
	err := launchCmd(Path, "upgrade", app)
	if err != nil {
		Install(app)
	} else {
		fmt.Printf("App %s successfully updated\n", app)
	}
}

func GetPath() string {
	path, err := exec.LookPath(installerName)

	// If installer is installed all ok
	if err == nil {
		return path
	}

	// Try to install the installer
	fmt.Printf("I didn't find %s, which I need to install apps\n", installerName)
	fmt.Println("Attempting to install it....")
	if path, err = installItself(); err != nil {
		fmt.Println("Sorry, looks like you will need to install it manually")
		fmt.Println("I am opening for you the installer's home page.")
		fmt.Println("Please follow the instructions")
		exec.Command("open", installerHomePage).Run()
		os.Exit(1)
	}
	return path
}
