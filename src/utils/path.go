package utils

import (
	"fmt"
	"github.com/mitchellh/go-homedir"
	"os"
)

func Exist(filesPath []string) error {
	for _, file := range filesPath {
		if _, err := os.Stat(file); os.IsNotExist(err) {
			return fmt.Errorf("file %s does not exist", file)
		}
	}
	return nil
}

func GetDefaultDownloadFolder() (string, error) {
	// Determine default download folder
	homeDir, err := homedir.Dir()
	if err != nil {
		return "", err
	}
	return homedir.Expand(homeDir + "/Downloads")
}
