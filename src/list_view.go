package src

import (
	"fmt"
	"os"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var (
	titleStyle      = lipgloss.NewStyle().MarginLeft(2)
	paginationStyle = list.DefaultStyles().PaginationStyle.PaddingLeft(4)
	helpStyle       = list.DefaultStyles().HelpStyle.PaddingLeft(4).PaddingBottom(1)
)

type ListItem struct {
	T, D string
}

func (i ListItem) Title() string       { return i.T }
func (i ListItem) Description() string { return i.D }
func (i ListItem) FilterValue() string { return i.T }

type ListViewModel struct {
	list     list.Model
	selected string
	endValue *string
	quitting bool
}

func (m ListViewModel) Init() tea.Cmd {
	return nil
}

func (m ListViewModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.list.SetWidth(msg.Width)
		return m, nil

	case tea.KeyMsg:
		switch keypress := msg.String(); keypress {
		case "q", "ctrl+c":
			m.quitting = true
			return m, tea.Quit

		case "enter":
			m.quitting = true
			i, ok := m.list.SelectedItem().(ListItem)
			if ok {
				*m.endValue = string(i.T)
			}
			return m, tea.Quit
		}
	}

	i, ok := m.list.SelectedItem().(ListItem)
	if ok {
		m.selected = string(i.D)
	}

	var cmd tea.Cmd
	m.list, cmd = m.list.Update(msg)
	return m, cmd
}

func (m ListViewModel) View() string {
	if m.quitting {
		return ""
	}

	return m.list.View()
}

func ListView(title string, op []ListItem, height int, endValue *string) {

	items := []list.Item{}
	for _, o := range op {
		items = append(items, o)
	}

	const defaultWidth = 20

	l := list.New(items, list.NewDefaultDelegate(), defaultWidth, height)
	l.Title = title
	l.SetShowStatusBar(false)
	l.SetFilteringEnabled(true)
	l.Styles.Title = titleStyle
	l.Styles.PaginationStyle = paginationStyle
	l.Styles.HelpStyle = helpStyle

	m := ListViewModel{list: l, endValue: endValue, selected: ""}

	if _, err := tea.NewProgram(m).Run(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}
}
