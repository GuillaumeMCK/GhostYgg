package utils

import (
	"fmt"
	"time"
)

const (
	UP   = "\033[A"
	DOWN = "\033[B"
)

// withUnit default value is true
func ByteSuffixes(i int64, withUnit ...bool) string {

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

	if len(withUnit) > 0 && !withUnit[0] {
		return fmt.Sprintf("%.1f", value)
	} else {
		return fmt.Sprintf("%.1f%s", value, suffix)
	}
}

func GetDateTime() string {
	return time.Now().Format("02/01/2006 15:04:05")
}
