package src

import (
	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/lipgloss"
)

type Styles struct {
	FooterColor        lipgloss.Color
	BorderColor        lipgloss.Color
	TitleColor         lipgloss.Color
	SelectedTitleColor lipgloss.Color

	FooterStyle         lipgloss.Style
	TitleStyle          lipgloss.Style
	InputField          lipgloss.Style
	InputFieldWithError lipgloss.Style

	PaginationStyle   lipgloss.Style
	HelpStyle         lipgloss.Style
	SelectedItemStyle lipgloss.Style

	PeachColor      lipgloss.Color
	CoralColor      lipgloss.Color
	OrchidColor     lipgloss.Color
	ThistleColor    lipgloss.Color
	NyanzaColor     lipgloss.Color
	AquamarineColor lipgloss.Color
	ErrorColor      lipgloss.Color
}

func DefaultStyles() *Styles {
	s := new(Styles)

	s.PeachColor = lipgloss.Color("#F2B391")
	s.CoralColor = lipgloss.Color("#F39194")
	s.OrchidColor = lipgloss.Color("#E3B5BF")
	s.ThistleColor = lipgloss.Color("#DAC3E9")
	s.NyanzaColor = lipgloss.Color("#E9F2D0")
	s.ErrorColor = lipgloss.Color("#FF99B8")
	s.AquamarineColor = lipgloss.Color("#B4F8D5")

	s.BorderColor = s.OrchidColor
	s.FooterColor = s.NyanzaColor
	s.TitleColor = s.ThistleColor
	s.SelectedTitleColor = s.OrchidColor

	s.InputField = lipgloss.NewStyle().BorderForeground(s.BorderColor).BorderStyle(lipgloss.NormalBorder()).Padding(1).Width(80)
	s.InputFieldWithError = lipgloss.NewStyle().BorderForeground(s.ErrorColor).BorderStyle(lipgloss.NormalBorder()).Padding(1).Width(80)
	s.FooterStyle = lipgloss.NewStyle().PaddingLeft(1).Foreground(s.FooterColor).Italic(true)
	s.TitleStyle = lipgloss.NewStyle().PaddingLeft(1).Foreground(s.TitleColor).Bold(true)

	s.PaginationStyle = list.DefaultStyles().PaginationStyle.PaddingLeft(4)
	s.HelpStyle = list.DefaultStyles().HelpStyle.PaddingLeft(4).PaddingBottom(1)
	s.SelectedItemStyle = lipgloss.NewStyle().
		Border(lipgloss.NormalBorder(), false, false, false, true).
		BorderForeground(s.SelectedTitleColor).
		Foreground(s.SelectedTitleColor).
		Padding(0, 0, 0, 1)

	return s
}

func (s Styles) Text(t string, c lipgloss.Color) string {
	var style = lipgloss.NewStyle().Foreground(c)
	return style.Render(t)
}
