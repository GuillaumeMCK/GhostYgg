package utils

import (
	"fmt"
	"time"
)

const (
	UP   = "\033[A"
	DOWN = "\033[B"
)

const (
	_  = iota             // ignore first value by assigning to blank identifier
	KB = 1 << (10 * iota) // 1024
	MB
	GB
	TB
)

func byteSuffixes(i int64) (suffix string, unit float64) {
	switch {
	case i < KB:
		suffix = "B"
		unit = 1
	case i < MB:
		suffix = "KB"
		unit = KB
	case i < GB:
		suffix = "MB"
		unit = MB
	case i < TB:
		suffix = "GB"
		unit = GB
	default:
		suffix = "TB"
		unit = TB
	}
	return
}

func FormatBytesProgress(bytesCompleted, totalLength int64) string {
	suffix, unit := byteSuffixes(totalLength)
	return fmt.Sprintf("%.1f/%.1f%s",
		float64(bytesCompleted)/unit,
		float64(totalLength)/unit,
		suffix)
}

func GetDateTime() string {
	return time.Now().Format("02/01/2006 15:04:05")
}

func ClearScreen() {
	// Clear the screen by printing ANSI escape sequences
	fmt.Print("\033[H\033[2J")
	// Move the cursor to the top left
	fmt.Print("\033[1;1H")
}
