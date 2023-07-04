package utils

import (
	"fmt"
	"time"
)

const (
	UP   = "\033[A"
	DOWN = "\033[B"
)

func ByteSuffixes(i int64) string {
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

func GetDateTime() string {
	return time.Now().Format("02/01/2006 15:04:05")
}