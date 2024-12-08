package src

import (
	"fmt"
	"os"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type Styles struct {
	BorderColor lipgloss.Color
	InputField  lipgloss.Style
}

func DefaultStyles() *Styles {
	s := new(Styles)
	s.BorderColor = lipgloss.Color("36")
	s.InputField = lipgloss.NewStyle().BorderForeground(s.BorderColor).BorderStyle(lipgloss.NormalBorder()).Padding(1).Width(80)
	return s
}

type (
	errMsg error
)

type textInputViewModel struct {
	textInput textinput.Model
	err       error
	question  string
	endValue  *string
	quitting  bool
	styles    *Styles
}

func TextFieldViewModel(question, placeHolder string, value *string) textInputViewModel {
	ti := textinput.New()
	ti.Placeholder = placeHolder
	ti.Focus()
	ti.CharLimit = 156
	ti.Width = 28
	ti.Placeholder = placeHolder

	return textInputViewModel{
		textInput: ti,
		err:       nil,
		question:  question,
		endValue:  value,
		styles:    DefaultStyles(),
		quitting:  false,
	}
}

func (m textInputViewModel) Init() tea.Cmd {
	return textinput.Blink
}

func (m textInputViewModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyEnter, tea.KeyCtrlC, tea.KeyEsc:
			*m.endValue = m.textInput.Value()
			m.quitting = true
			return m, tea.Quit
		}

	case errMsg:
		m.err = msg
		return m, nil
	}

	m.textInput, cmd = m.textInput.Update(msg)
	return m, cmd
}

func (m textInputViewModel) View() string {
	if m.quitting {
		return ""
	}

	return lipgloss.JoinVertical(
		lipgloss.Left,
		fmt.Sprintf("\n%s\n", m.question),
		m.styles.InputField.Render(m.textInput.View()),
		"\n(ctrl+c to quit)",
	)
}

func TextFieldView(title, placeHolder string, endValue *string) {

	m := TextFieldViewModel(title, placeHolder, endValue)

	if _, err := tea.NewProgram(m).Run(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}
}
