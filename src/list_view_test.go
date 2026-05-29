package src

import (
	"testing"

	tea "github.com/charmbracelet/bubbletea"
)

func TestListViewModelClearsViewOnCancel(t *testing.T) {
	tests := []struct {
		name string
		msg  tea.KeyMsg
	}{
		{"esc", tea.KeyMsg{Type: tea.KeyEsc}},
		{"ctrl+c", tea.KeyMsg{Type: tea.KeyCtrlC}},
		{"q", tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune{'q'}}},
	}

	for _, test := range tests {
		endValue := ListItem{}
		model := ListViewModel{endValue: &endValue}

		updatedModel, cmd := model.Update(test.msg)

		if cmd == nil {
			t.Errorf("expected quit command for key %q", test.name)
		}

		if endValue.T != ExitSignal {
			t.Errorf("expected exit signal for key %q, got %q", test.name, endValue.T)
		}

		if view := updatedModel.View(); view != "" {
			t.Errorf("expected empty view after cancel for key %q, got %q", test.name, view)
		}
	}
}
