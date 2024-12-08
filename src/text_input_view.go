package src

import (
	"fmt"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

type (
	errMsg error
)

type textInputViewModel struct {
	textInput textinput.Model
	err       error
	question  string
	endValue  *string
	quitting  bool
}

func TextInputView(question, placeHolder string, value *string) textInputViewModel {
	ti := textinput.New()
	ti.Placeholder = placeHolder
	ti.Focus()
	ti.CharLimit = 156
	ti.Width = 20

	return textInputViewModel{
		textInput: ti,
		err:       nil,
		question:  question,
		endValue:  value,
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

	return fmt.Sprintf(
		"%s\n\n%s\n\n%s",
		m.question,
		m.textInput.View(),
		"(esc to quit)",
	) + "\n"
}
