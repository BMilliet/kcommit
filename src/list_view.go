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
}

func ListView(t string, li []ListItem, v *string) ListViewModel {
	tl := t + "\n"

	return ListViewModel{
		title:    tl,
		choices:  li,
		endValue: v,
	}
}

func (m ListViewModel) Init() tea.Cmd {
	// Just return `nil`, which means "no I/O right now, please."
	return nil
}

func (m ListViewModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	// Is it a key press?
	case tea.KeyMsg:

		// Cool, what was the actual key pressed?
		switch msg.String() {

		// These keys should exit the program.
		case "ctrl+c", "q":
			return m, tea.Quit

		// The "up" and "k" keys move the cursor up
		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			}

		// The "down" and "j" keys move the cursor down
		case "down", "j":
			if m.cursor < len(m.choices)-1 {
				m.cursor++
			}

		case "enter", " ":
			*m.endValue = m.choices[m.cursor].Title
			return m, tea.Quit
		}
	}

	// Return the updated ListViewModel to the Bubble Tea runtime for processing.
	// Note that we're not returning a command.
	return m, nil
}

func (m ListViewModel) View() string {

	// Iterate over our choices
	for i, choice := range m.choices {

		// Is the cursor pointing at this choice?
		cursor := " " // no cursor
		if m.cursor == i {
			cursor = "->"
		}

		// Render the row
		m.title += fmt.Sprintf("%s %s\n", cursor, choice.Title)
	}

	// The footer
	m.title += "\nPress q to quit.\n"

	// Send the UI for rendering
	return m.title
}
