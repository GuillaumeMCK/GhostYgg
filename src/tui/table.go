package tui

import (
	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"math"
	"strings"
)

// TableModel represents the model for the table.
type TableModel struct {
	table table.Model
	ctx   *TableCtx
	exit  bool
}

// TableCtx represents the configuration for the table.
type TableCtx struct {
	Columns []string
	Rows    [][]string
	Widths  []float32
}

var baseStyle = lipgloss.NewStyle().
	BorderStyle(lipgloss.RoundedBorder()).
	BorderForeground(lipgloss.Color("240"))

type UpdateTableMsg struct{}

func doUpdateTable() tea.Cmd {
	return func() tea.Msg {
		return UpdateTableMsg{}
	}
}

func (m TableModel) Init() tea.Cmd {
	m.table = m.resize()   // Initialize the table.
	return doUpdateTable() // Update the table.
}

// Update handles the update messages for the table model.
func (m TableModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
	case UpdateTableMsg:
		m.table = m.resize()
		return m, doUpdateTable()
	case tea.KeyMsg:
		switch msg.String() {
		case "esc", "q", "ctrl+c":
			m.exit = true
			return m, tea.Quit
		case "enter":
			break
		}
	}
	m.table, cmd = m.table.Update(msg)
	return m, cmd
}

// View returns the view for the table model.
func (m TableModel) View() string {
	return baseStyle.Render(m.table.View())
}

// Exit returns true if the table should exit.
func (m TableModel) Exit() bool {
	return m.exit
}

// resize resizes the table.
func (m TableModel) resize() table.Model {
	rows, columns, height := generateTableContent(m.ctx)
	m.table.SetColumns(columns)
	m.table.SetRows(rows)
	m.table.SetHeight(height)
	return m.table
}

// Table creates a new table.
func Table(ctx *TableCtx) TableModel {
	rows, columns, height := generateTableContent(ctx)
	t := table.New(
		table.WithColumns(columns),
		table.WithRows(rows),
		table.WithFocused(true),
		table.WithHeight(height),
	)
	s := table.DefaultStyles()
	s.Header = s.Header.
		BorderStyle(lipgloss.NormalBorder()).
		BorderForeground(lipgloss.Color("240")).
		BorderBottom(true).
		Bold(true)
	s.Selected = s.Selected.
		Foreground(lipgloss.Color("229")).
		Background(lipgloss.Color("57")).
		Bold(false)
	t.SetStyles(s)
	return TableModel{table: t, ctx: ctx, exit: false}
}

// generateTable generates a table based on the input rows and columns.
func generateTableContent(ctx *TableCtx) ([]table.Row, []table.Column, int) {
	screenWidth, screenHeight, _ := GetTerminalSize()
	//println("screenWidth", screenWidth, "screenHeight", screenHeight)
	//println("ctx.Columns", len(ctx.Columns), "ctx.Widths", len(ctx.Widths), "ctx.Rows", len(ctx.Rows))
	columns := createColumns(ctx.Columns, ctx.Widths, screenWidth)
	rows := createRows(ctx.Rows)
	height := screenHeight - 4
	width := screenWidth - 4
	for _, column := range columns {
		width -= column.Width + 1 // Add 1 for the column separator
	}
	if width > 0 {
		for i := range columns {
			columns[i].Width += width / len(columns) // Distribute the remaining width evenly among columns
			width -= width / len(columns)            // Update the remaining width
			if width <= 0 {
				break
			}
		}
	}
	return rows, columns, height
}

// createColumns creates table columns based on column names and widths.
func createColumns(columnNames []string, widthsPct []float32, screenWidth int) []table.Column {
	columns := make([]table.Column, 0)
	totalColumnWidth := 0
	for _, pct := range widthsPct {
		totalColumnWidth += int(math.Ceil(float64(screenWidth) * float64(pct)))
	}
	maxScreenSize := screenWidth - 2*(len(columnNames)+1)
	for i, pct := range widthsPct {
		w := int(math.Ceil(float64(screenWidth) * float64(pct)))
		if totalColumnWidth > maxScreenSize {
			w -= int(math.Ceil(float64(totalColumnWidth-maxScreenSize) / float64(len(widthsPct))))
		}
		columns = append(columns, table.Column{Title: columnNames[i], Width: w})
	}
	return columns
}

// createRows creates table rows based on the input rows.
func createRows(rowsString [][]string) []table.Row {
	rows := make([]table.Row, 0)
	for _, row := range rowsString {
		for i, el := range row {
			row[i] = strings.TrimSpace(el)
		}
		rows = append(rows, row)
	}
	return rows
}
