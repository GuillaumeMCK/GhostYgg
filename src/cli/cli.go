package cli

import (
	"fmt"
)

const (
	UP           = "\033[A"    // ANSI escape sequence to move the cursor up one line
	DOWN         = "\033[B"    // ANSI escape sequence to move the cursor down one line
	RESET        = "\033[0m"   // ANSI escape sequence to reset all attributes
	CURSOR_START = "\033[1;1H" // ANSI escape sequence to move the cursor to the start of the line
	CLEAR_SCREEN = "\033[2J"   // ANSI escape sequence to clear the screen
)

func ClearScreen() {
	fmt.Print(CLEAR_SCREEN)
	fmt.Print(CURSOR_START)
}
