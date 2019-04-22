package utils

import (
	"currencyRate/config"
	"os"
	"strings"
)

// SetDirectoryTree creates directory tree
func SetDirectoryTree() {
	path := config.OutputPath
	bits := strings.Split(path, "/")
	appendedPath := ""

	for _, bit := range bits {
		if bit != "" {
			appendedPath = appendedPath + "/" + bit
			if _, err := os.Stat(appendedPath); os.IsNotExist(err) {
				os.Mkdir(appendedPath, os.ModePerm)
			}
		}
	}

}
