package utils

import (
	"fmt"
	"github.com/mitchellh/go-homedir"
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

func GetDefaultDownloadFolder() (string, error) {
	homeDir, err := homedir.Dir()
	if err != nil {
		return "", err
	}

	downloadFolder := ""
	switch {
	case os.Getenv("HOME") != "":
		downloadFolder = os.Getenv("HOME") + "/Downloads"
	case os.Getenv("USERPROFILE") != "":
		downloadFolder = os.Getenv("USERPROFILE") + "\\Downloads"
	default:
		downloadFolder = homeDir
	}

	return downloadFolder, nil
}
