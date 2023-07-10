package utils

import (
	"GhostYgg/src/cli"
	"bytes"
	"fmt"
	"golang.org/x/crypto/ssh/terminal"
	"os"
	"regexp"
	"strings"
	"time"
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

// PrintRow prints a string on a specific row of the terminal, adjusting the row position and truncating the string if necessary.
func PrintRow(rowIndex int, s string) {
	down := strings.Repeat(cli.DOWN, rowIndex)
	up := strings.Repeat(cli.UP, rowIndex)

	consoleWidth, _, err := terminal.GetSize(int(os.Stdout.Fd())) // get the width of the console in characters
	if err != nil {
		fmt.Println(err)
		return
	}
	s = truncate(s, consoleWidth-1) // truncate the string if its length exceeds the width of the console

	fmt.Printf("%s%s\r%s%s", cli.RESET, down, s, up)
}

// truncate truncates the given string if its length exceeds the maximum limit, appending ellipsis.
func truncate(s string, width int) string {
	printableLen := len(s) - countEscapeCharsAndColors(s)
	if printableLen > width {
		// Calculate the number of characters needed for ellipsis
		ellipsisWidth := 3
		// Create a buffer to build the truncated string
		var truncated bytes.Buffer
		truncated.Grow(width + ellipsisWidth)

		// Iterate over the string and truncate as necessary
		for _, char := range s {
			truncated.WriteRune(char)
			if truncated.Len()-countEscapeCharsAndColors(truncated.String()) >= width {
				break
			}
		}
		// Remove the last three characters and append ellipsis
		truncated.Truncate(truncated.Len() - ellipsisWidth)
		truncated.WriteString("...")
		return truncated.String()
	}

	// Calculate the spacing needed to reach the specified width
	spacing := width - printableLen
	// Append the required spacing to the original string
	return s + strings.Repeat(" ", spacing)
}

// countEscapeCharsAndColors counts the number of ANSI escape characters and color codes in the given string.
func countEscapeCharsAndColors(input string) int {
	// improve regex. add reset color code and other special characters that not shown in the terminal
	re := regexp.MustCompile(`\033\[[0-9;]*[a-zA-Z]`)
	// find all matches in the string
	matches := re.FindAllString(input, -1)
	return len(matches) * 5 // each escape character ≈ 4 characters
}
