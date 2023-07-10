package cli

import (
	"fmt"
	"io"
	"os"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

const (
	listHeight   = 14
	defaultWidth = 20
)

var (
	titleStyle        = lipgloss.NewStyle().MarginLeft(2)
	itemStyle         = lipgloss.NewStyle().PaddingLeft(4)
	selectedItemStyle = lipgloss.NewStyle().PaddingLeft(2).Foreground(lipgloss.Color("170"))
	paginationStyle   = list.DefaultStyles().PaginationStyle.PaddingLeft(4)
	quitTextStyle     = lipgloss.NewStyle().Margin(1, 0, 2, 4)
)

// Item represents a list item.
type Item string

// FilterValue returns the filter value for an item.
func (i Item) FilterValue() string { return "" }

// itemDelegate implements the list.ItemDelegate interface.
type itemDelegate struct{}

// Height returns the height of the list item.
func (d itemDelegate) Height() int { return 1 }

// Spacing returns the spacing between list items.
func (d itemDelegate) Spacing() int { return 0 }

// Update handles the update message for a list item.
func (d itemDelegate) Update(msg tea.Msg, _ *list.Model) tea.Cmd { return nil }

// Render renders a list item.
func (d itemDelegate) Render(w io.Writer, m list.Model, index int, listItem list.Item) {
	i, ok := listItem.(Item)
	if !ok {
		return
	}

	str := fmt.Sprintf("%s", i)

	fn := itemStyle.Render
	if index == m.Index() {
		fn = selectedItemStyle.Render
	}

	fmt.Fprint(w, fn(str))
}

// ListModel represents the model for the list.
type ListModel struct {
	list     list.Model
	choice   string
	quitting bool
}

// Init initializes the list model.
func (m ListModel) Init() tea.Cmd {
	return nil
}

// Update handles the update messages for the list model.
func (m ListModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.list.SetWidth(msg.Width)
		return m, nil

	case tea.KeyMsg:
		switch keypress := msg.String(); keypress {
		case "ctrl+c":
			m.quitting = true
			return m, tea.Quit

		case "enter":
			i, ok := m.list.SelectedItem().(Item)
			if ok {
				m.choice = string(i)
			}
			return m, tea.Quit
		}
	}

	var cmd tea.Cmd
	m.list, cmd = m.list.Update(msg)
	return m, cmd
}

// View returns the view for the list model.
func (m ListModel) View() string {
	if m.choice != "" {
		return m.choice
	}
	return "\n" + m.list.View()
}

var choiceMap = map[string]int{
	"Show credits": 0,
	"Show reviews": 1,
	"Go back":      2,
}

// PrintList displays the list and returns the selected choice index.
func PrintList() int {
	items := []list.Item{Item("Show credits"), Item("Show reviews"), Item("Go back")}

	l := list.New(items, itemDelegate{}, defaultWidth, listHeight)
	l.Title = "Select operation:"
	l.SetShowStatusBar(false)
	l.SetFilteringEnabled(false)
	l.Styles.Title = titleStyle
	l.Styles.PaginationStyle = paginationStyle
	l.SetShowHelp(false)

	p := tea.NewProgram(ListModel{list: l})
	m, err := p.Run()
	if err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}
	if m, ok := m.(ListModel); ok {
		if !m.quitting {
			return choiceMap[m.choice]
		}
	} else {
		fmt.Println("Error in table")
		os.Exit(1)
	}

	return -1
}
