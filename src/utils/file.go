package utils

import (
	"fmt"
	"os"
	"path/filepath"
)

func Exist(downloadFolder string, file string) error {
	// Check if the file exists in the download folder
	filePath := filepath.Join(downloadFolder, file)
	if _, err := os.Stat(filePath); err == nil {
		return fmt.Errorf("le fichier %s existe déjà dans le répertoire de téléchargement", file)
	}

	return nil
}
