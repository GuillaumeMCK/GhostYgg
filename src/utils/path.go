package utils

import (
	"github.com/mitchellh/go-homedir"
	"os"
)

func Exist(filesPath string) bool {
	_, err := os.Stat(filesPath)
	return err == nil
}

func GetDefaultDownloadFolder() (string, error) {
	homeDir, err := GetHomeDir()
	if err != nil {
		return "", err
	}
	return homedir.Expand(homeDir + "/Downloads")
}

func GetHomeDir() (string, error) {
	return homedir.Dir()
}
