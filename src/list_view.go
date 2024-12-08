package src

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
)

type ListItem struct {
	Title string
	Desc  string
}

type ListViewModel struct {
	title    string
	choices  []ListItem
	cursor   int
	endValue *string
	quitting bool
}

func ListView(t string, li []ListItem, v *string) ListViewModel {
	tl := t + "\n"

	return ListViewModel{
		title:    tl,
		choices:  li,
		endValue: v,
		quitting: false,
	}
}

func (m ListViewModel) Init() tea.Cmd {
	return nil
}

func (m ListViewModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case tea.KeyMsg:

		switch msg.String() {

		case "ctrl+c", "q":
			m.quitting = true
			return m, tea.Quit

		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			}

		case "down", "j":
			if m.cursor < len(m.choices)-1 {
				m.cursor++
			}

		case "enter", " ":
			*m.endValue = m.choices[m.cursor].Title
			m.quitting = true
			return m, tea.Quit
		}
	}

	return m, nil
}

func (m ListViewModel) View() string {

	if m.quitting {
		return ""
	}

	for i, choice := range m.choices {

		cursor := " "
		if m.cursor == i {
			cursor = "->"
		}

		m.title += fmt.Sprintf("%s %s\n", cursor, choice.Title)
	}

	m.title += "\nPress q to quit.\n"

	return m.title
}
