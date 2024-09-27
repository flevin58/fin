package tools

import (
	"os"
	"path"
	"strings"
)

func NormalizePath(filePath string, root string) string {
	if strings.HasPrefix(filePath, "~/") {
		filePath = path.Join("$HOME", filePath[2:])
	}
	filePath = os.ExpandEnv(filePath)
	if root != "" && !path.IsAbs(filePath) {
		filePath = os.ExpandEnv(path.Join(root, filePath))
	}
	return filePath
}
