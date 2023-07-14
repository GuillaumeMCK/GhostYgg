package table

import (
	"GhostYgg/src/tui/constants"
	tab "github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
	"math"
)

// Model represents the model of the table.
type Model struct {
	table tab.Model
	ctx   *TableCtx
}

// TableCtx represents the context of the table. It contains the columns, rows and widths.
type TableCtx struct {
	Columns [6]string
	Rows    [][6]string
	Widths  [6]float32
}

// Init initializes the table on creation.
func (m Model) Init() tea.Cmd {
	m.refresh()
	return nil
}

// Update handles the update messages for the model.
func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch _ := msg.(type) {
	case tea.WindowSizeMsg:
	case SelectedRowMsg:
		return m, selectedRow(msg.(SelectedRowMsg).Index)
	case UpdateTableMsg:
		m.refresh()
		return m, forceUpdateTable()
	}
	m.table, cmd = m.table.Update(msg)
	return m, cmd
}

// View returns the view for the model.
func (m Model) View() string {
	return constants.BaseTableStyle.Render(m.table.View())
}

// refresh refreshes the table.
func (m *Model) refresh() {
	rows, columns, height := generateTableContent(m.ctx)
	m.table.SetColumns(columns)
	m.table.SetRows(rows)
	m.table.SetHeight(height)
}

// New creates a new table based on the input context.
func New(ctx *TableCtx) Model {
	rows, columns, height := generateTableContent(ctx)
	t := tab.New(
		tab.WithColumns(columns),
		tab.WithRows(rows),
		tab.WithFocused(true),
		tab.WithHeight(height),
	)
	t.SetStyles(constants.TableStyle)
	return Model{table: t, ctx: ctx}
}

// generateTableContent generates the content of the table based on the input context.
func generateTableContent(ctx *TableCtx) ([]tab.Row, []tab.Column, int) {
	screenWidth, screenHeight := constants.WindowSize.Width, constants.WindowSize.Height
	columns := createColumns(ctx.Columns, ctx.Widths, screenWidth)
	rows := createRows(ctx.Rows)
	height := screenHeight - 4 - constants.HelpHeight
	width := screenWidth - 4
	for _, column := range columns {
		width -= column.Width + 12 // Add 1 for the column separator
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

// createColumns creates columns based on column names and widths.
func createColumns(columnNames [6]string, widths [6]float32, screenWidth int) []tab.Column {
	columns := make([]tab.Column, 0)
	totalColumnWidth := 0
	for _, pct := range widths {
		totalColumnWidth += int(math.Ceil(float64(screenWidth) * float64(pct)))
	}
	maxScreenSize := screenWidth - 2*(len(columnNames)+1)
	for i, pct := range widths {
		w := int(math.Ceil(float64(screenWidth) * float64(pct)))
		if totalColumnWidth > maxScreenSize {
			w -= int(math.Ceil(float64(totalColumnWidth-maxScreenSize) / float64(len(widths))))
		}
		columns = append(columns, tab.Column{Title: columnNames[i], Width: w})
	}
	return columns
}

// createRows creates rows based on the input rows.
func createRows(rowsString [][6]string) []tab.Row {
	rows := make([]tab.Row, 0)
	for _, row := range rowsString {
		rows = append(rows, row[:])
	}
	return rows
}
