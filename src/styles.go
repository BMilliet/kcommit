package src

import (
	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/lipgloss"
)

var (
	primaryColor   = lipgloss.Color("#FF77FF")
	secondaryColor = lipgloss.Color("#FF99CC")
	thirdColor     = lipgloss.Color("#CC99FF")
	logColor       = lipgloss.Color("#57cc99")
	errorColor     = lipgloss.Color("#D1007A")
)

var (
	FooterMsgStyle = lipgloss.NewStyle().PaddingLeft(0).Foreground(logColor).Italic(true)
	LogMsgStyle    = lipgloss.NewStyle().PaddingLeft(1).Foreground(logColor).Italic(true)
	ErrorMsgStyle  = lipgloss.NewStyle().PaddingLeft(1).Foreground(errorColor).Bold(true)

	FocusedStyle          = lipgloss.NewStyle().Foreground(thirdColor).Bold(true)
	TitleStyle            = lipgloss.NewStyle().Foreground(thirdColor).Bold(true).Padding(0, 1, 0)
	SelectedItemDescStyle = lipgloss.NewStyle().PaddingLeft(1).Foreground(secondaryColor)
	DescriptionStyle      = lipgloss.NewStyle().Foreground(secondaryColor)

	ItemStyle         = lipgloss.NewStyle().PaddingLeft(4)
	SelectedItemStyle = lipgloss.NewStyle().PaddingLeft(2).Foreground(primaryColor)
	PaginationStyle   = list.DefaultStyles().PaginationStyle.PaddingLeft(4)
	HelpStyle         = list.DefaultStyles().HelpStyle.PaddingLeft(4).PaddingBottom(1)
	QuitTextStyle     = lipgloss.NewStyle().Margin(1, 0, 2, 4)

	BorderColor = thirdColor
	InputField  = lipgloss.NewStyle().BorderForeground(
		BorderColor,
	).BorderStyle(
		lipgloss.NormalBorder(),
	).Padding(1).Width(80)
)
