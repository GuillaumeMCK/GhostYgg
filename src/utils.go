package main

import (
	"fmt"
	"os"
	"path/filepath"
	"time"
)

const (
	UP   = "\033[A"
	DOWN = "\033[B"
)

func byteSuffixes(i int64) string {
	const (
		_  = iota             // ignore first value by assigning to blank identifier
		KB = 1 << (10 * iota) // 1024
		MB
		GB
		TB
	)

	var suffix string
	value := float64(i)

	switch {
	case i < KB:
		suffix = "B"
	case i < MB:
		suffix = "KB"
		value /= KB
	case i < GB:
		suffix = "MB"
		value /= MB
	case i < TB:
		suffix = "GB"
		value /= GB
	default:
		suffix = "TB"
		value /= TB
	}

	return fmt.Sprintf("%.1f%s", value, suffix)
}

func getDateTime() string {
	return time.Now().Format("02/01/2006 15:04:05")
}

func checkFileExistence(downloadFolder string, file string) error {
	// Check if the file exists in the download folder
	filePath := filepath.Join(downloadFolder, file)
	if _, err := os.Stat(filePath); err == nil {
		return fmt.Errorf("le fichier %s existe déjà dans le répertoire de téléchargement", file)
	}

	return nil
}
