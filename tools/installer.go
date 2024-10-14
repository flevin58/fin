package tools

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"
)

var (
	InstallerName = installerName
	InstallerPath = GetPath()
)

// Launches a command and returns the stdout as a slice of lines.
// If verbose is false, no output is written to the console
func launchCmd(verbose bool, cmd string, args ...string) ([]string, error) {
	var buffer bytes.Buffer
	var output io.Writer
	if verbose {
		output = io.MultiWriter(&buffer, os.Stdout)
	} else {
		output = &buffer
	}
	command := exec.Command(cmd, args...)
	command.Stdin = os.Stdin
	command.Stdout = output
	command.Stderr = os.Stderr
	err := command.Run()
	lines := strings.Split(buffer.String(), "\n")
	return lines, err
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

func List() ([]string, error) {
	lines, err := launchCmd(false, InstallerPath, "leaves")
	if err != nil {
		return []string{}, err
	}
	return lines, nil
}
