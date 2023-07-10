package main

import (
	"GhostYgg/src/cli"
	"fmt"
)

func main() {
	columns := []string{"Name", "Age", "City"}
	rows := [][]string{
		{"John Doe", "25", "New York"},
		{"Jane Smith", "30", "San Francisco"},
		{"Bob Johnson", "40", "Chicago"},
	}

	// Set the desired widths for each column
	widths := []int{20, 10, 15}

	// Print the table and get the selected row index
	selectedRow := cli.PrintRows(rows, columns, widths)

	if selectedRow >= 0 {
		selectedData := rows[selectedRow]
		fmt.Printf("Selected row: %v\n", selectedData)
	} else {
		fmt.Println("No row selected.")
	}

	// Wait for user input before exiting
	fmt.Println("Press Enter to exit...")
	_, _ = fmt.Scanln()
}
