package src

import (
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

const listHeight = 14

var (
	titleStyle        = lipgloss.NewStyle().MarginLeft(2)
	itemStyle         = lipgloss.NewStyle().PaddingLeft(4)
	selectedItemStyle = lipgloss.NewStyle().PaddingLeft(2).Foreground(lipgloss.Color("170"))
	paginationStyle   = list.DefaultStyles().PaginationStyle.PaddingLeft(4)
	helpStyle         = list.DefaultStyles().HelpStyle.PaddingLeft(4).PaddingBottom(1)
)

type itemDelegate struct{}

func (d itemDelegate) Height() int                             { return 1 }
func (d itemDelegate) Spacing() int                            { return 0 }
func (d itemDelegate) Update(_ tea.Msg, _ *list.Model) tea.Cmd { return nil }
func (d itemDelegate) Render(w io.Writer, m list.Model, index int, listItem list.Item) {
	i, ok := listItem.(ListItem)
	if !ok {
		return
	}

	str := fmt.Sprintf("%d. %s", index+1, i.T)

	fn := itemStyle.Render
	if index == m.Index() {
		fn = func(s ...string) string {
			return selectedItemStyle.Render("> " + strings.Join(s, " "))
		}
	}

	fmt.Fprint(w, fn(str))
}

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

	if m.selected != "" {
		return m.list.View() + m.selected
	}

	return m.list.View()
}

func ListView(title string, op []ListItem, endValue *string) {

	items := []list.Item{}
	for _, o := range op {
		items = append(items, o)
	}

	const defaultWidth = 20

	l := list.New(items, itemDelegate{}, defaultWidth, listHeight)
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
