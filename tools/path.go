package tools

import (
	"os"
	"path"
	"strings"
)

func IsValidFolder(path string) bool {
	file, err := os.Open(path)
	if err != nil {
		return false
	}
	defer file.Close()
	fileInfo, err := file.Stat()
	if err != nil {
		return false
	}
	return fileInfo.IsDir()
}

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

func NormalizePathWithExt(filePath, rootPath, ext string) string {
	filePath = NormalizePath(filePath, rootPath)
	if path.Ext(filePath) == "" {
		filePath += ext
	}
	return filePath
}
