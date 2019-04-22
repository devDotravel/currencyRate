package utils

import (
	"currency/config"
	"os"
	"strings"
)

func SetDirectoryTree() {
	// creates availability folder and parent output root folder
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
