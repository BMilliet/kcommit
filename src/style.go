package src

import (
	"github.com/charmbracelet/lipgloss"
)

type Styles struct {
	FooterColor        lipgloss.Color
	BorderColor        lipgloss.Color
	TitleColor         lipgloss.Color
	SelectedTitleColor lipgloss.Color

	InputField  lipgloss.Style
	FooterStyle lipgloss.Style
	TitleStyle  lipgloss.Style

	PeachColor   lipgloss.Color
	CoralColor   lipgloss.Color
	OrchidColor  lipgloss.Color
	ThistleColor lipgloss.Color
	NyanzaColor  lipgloss.Color
}

func DefaultStyles() *Styles {
	s := new(Styles)

	s.PeachColor = lipgloss.Color("#F2B391")
	s.CoralColor = lipgloss.Color("#F39194")
	s.OrchidColor = lipgloss.Color("#E3B5BF")
	s.ThistleColor = lipgloss.Color("#DAC3E9")
	s.NyanzaColor = lipgloss.Color("#E9F2D0")

	s.BorderColor = s.OrchidColor
	s.FooterColor = s.NyanzaColor
	s.TitleColor = s.ThistleColor
	s.SelectedTitleColor = s.OrchidColor

	s.InputField = lipgloss.NewStyle().BorderForeground(s.BorderColor).BorderStyle(lipgloss.NormalBorder()).Padding(1).Width(80)
	s.FooterStyle = lipgloss.NewStyle().PaddingLeft(1).Foreground(s.FooterColor).Italic(true)
	s.TitleStyle = lipgloss.NewStyle().PaddingLeft(1).Foreground(s.TitleColor).Bold(true)
	return s
}

func (s Styles) Text(t string, c lipgloss.Color) string {
	var style = lipgloss.NewStyle().Foreground(c)
	return style.Render(t)
}
