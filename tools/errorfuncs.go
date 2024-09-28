package tools

import (
	"fmt"
	"os"
)

func Perror(format string, args ...any) {
	fmt.Printf("Error: ")
	fmt.Printf(format, args...)
	fmt.Println()
}

func ExitWithError(format string, args ...any) {
	Perror(format, args...)
	os.Exit(1)
}
