package utils

import (
	"fmt"
	"golang.org/x/crypto/ssh/terminal"
	"os"
	"regexp"
	"strings"
	"time"
)

const (
	UP   = "\033[A" // ANSI escape sequence to move the cursor up one line
	DOWN = "\033[B" // ANSI escape sequence to move the cursor down one line
)

const (
	_  = iota             // ignore first value by assigning to blank identifier
	KB = 1 << (10 * iota) // 1024
	MB
	GB
	TB
)

// byteSuffixes returns the appropriate suffix and unit based on the provided number of bytes.
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

// FormatBytesProgress formats the bytes completed and total length into a string representation with the appropriate suffix.
func FormatBytesProgress(bytesCompleted, totalLength int64) string {
	suffix, unit := byteSuffixes(totalLength)
	return fmt.Sprintf("%.1f/%.1f%s",
		float64(bytesCompleted)/unit,
		float64(totalLength)/unit,
		suffix)
}

// GetDateTime returns the current date and time formatted as "DD/MM/YYYY HH:MM:SS".
func GetDateTime() string {
	return time.Now().Format("02/01/2006 15:04:05")
}

// ClearScreen clears the terminal screen.
func ClearScreen() {
	// Clear the screen by printing ANSI escape sequences
	fmt.Print("\033[H\033[2J")
	// Move the cursor to the top left
	fmt.Print("\033[1;1H")
}

// truncate truncates the given string if its length exceeds the maximum limit, appending ellipsis.
func truncate(s string, max int) string {
	if len(s) > max {
		return s[:max-3] + "..."
	} else {
		spacing := max - (len(s) - countEscapeCharsAndColors(s))
		return s + fmt.Sprintf("%*s", spacing, " ")
	}
}

// countEscapeCharsAndColors counts the number of ANSI escape characters and color codes in the given string.
func countEscapeCharsAndColors(input string) int {
	re := regexp.MustCompile(`(\x1b\[[0-9;]+m)|\\[a-zA-Z]`)
	matches := re.FindAllString(input, -1)
	return len(matches)
}

// PrintRow prints a string on a specific row of the terminal, adjusting the row position and truncating the string if necessary.
func PrintRow(rowIndex int, s string) {
	down := strings.Repeat(DOWN, rowIndex)
	up := strings.Repeat(UP, rowIndex)

	consoleWidth, _, err := terminal.GetSize(int(os.Stdout.Fd()))
	if err != nil {
		fmt.Println(err)
	}

	consoleWidth += countEscapeCharsAndColors(s)

	s = truncate(s, consoleWidth)

	fmt.Printf("%s\r%s%s", down, s, up)
}
