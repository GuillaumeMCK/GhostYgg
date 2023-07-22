package tui

import (
	"GhostYgg/src/client"
	"GhostYgg/src/tui/constants"
	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
	"math"
)

// Table represents the model of the table.
type Table struct {
	table table.Model
	ctx   *TableCtx
}

// TableCtx represents the context of the table.
type TableCtx struct {
	Columns [6]string
	Rows    []client.TorrentInfos
	Widths  [6]float32
}

// Init initializes the table on creation.
func (m Table) Init() tea.Cmd {
	return nil
}

// Update handles the update messages for the model.
func (m Table) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	m.refresh()
	switch msg.(type) {
	case UpdateTableMsg:
		m.refresh()
		return m, cmd
	}
	m.table, cmd = m.table.Update(msg)
	return m, cmd
}

// View returns the view for the model.
func (m Table) View() string {
	return constants.BaseTableStyle.Render(m.table.View())
}

// refresh refreshes the table with new rows (if provided).
func (m *Table) refresh(newRows ...[]client.TorrentInfos) {
	if len(newRows) > 0 {
		m.ctx.Rows = newRows[0]
	}
	rows, columns, height := generateTableContent(m.ctx)
	m.table.SetColumns(columns)
	m.table.SetRows(rows)
	m.table.SetHeight(height)
}

// selectedRow returns the index of the selected row.
func (m Table) selectedRow() int {
	return m.table.Cursor()
}

// NewTable creates a new table based on the input context.
func NewTable(ctx *TableCtx) *Table {
	rows, columns, height := generateTableContent(ctx)
	t := table.New(
		table.WithColumns(columns),
		table.WithRows(rows),
		table.WithFocused(true),
		table.WithHeight(height),
	)
	t.SetStyles(constants.TableStyle)
	return &Table{table: t, ctx: ctx}
}

// generateTableContent generates the content of the table based on the input context.
func generateTableContent(ctx *TableCtx) ([]table.Row, []table.Column, int) {
	height := constants.WindowSize.Height - (4 + constants.HelpHeight)
	width := constants.WindowSize.Width - 4

	columns := createColumns(ctx.Columns, ctx.Widths, width)
	rows := createRows(ctx.Rows)
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

// createColumns creates columns based on column names and widths.
func createColumns(columnNames [6]string, widths [6]float32, screenWidth int) []table.Column {
	columns := make([]table.Column, 0)
	totalColumnWidth := 0
	for _, pct := range widths {
		totalColumnWidth += int(math.Ceil(float64(screenWidth) * float64(pct)))
	}
	maxScreenSize := screenWidth - 2*(len(columnNames)+1) // 2 for the padding on each side
	for i, pct := range widths {
		w := int(math.Ceil(float64(screenWidth) * float64(pct)))
		if totalColumnWidth > maxScreenSize {
			w -= int(math.Ceil(float64(totalColumnWidth-maxScreenSize) / float64(len(widths))))
		}
		columns = append(columns, table.Column{Title: columnNames[i], Width: w})
	}
	return columns
}

// createRows creates rows based on the input rows.
func createRows(Torrents []client.TorrentInfos) []table.Row {
	rows := make([]table.Row, 0)
	for _, row := range Torrents {
		rows = append(rows, table.Row{row.Infos.Name, row.Infos.Progress, row.Infos.Seeders, row.Infos.Leechers, row.Infos.Torrentspeed, row.Infos.ETA})
	}
	return rows
}
