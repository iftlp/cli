package app

import "github.com/charmbracelet/lipgloss"

type Styles struct {
	Title    lipgloss.Style
	Section  lipgloss.Style
	Label    lipgloss.Style
	Value    lipgloss.Style
	Error    lipgloss.Style
	QuitHint lipgloss.Style
}

func getStyles(theme string) Styles {
	if theme == "dark" {
		return Styles{
			Title:    lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("205")).MarginBottom(1),
			Section:  lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("69")).MarginTop(1).Underline(true),
			Label:    lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("212")),
			Value:    lipgloss.NewStyle().Foreground(lipgloss.Color("229")),
			Error:    lipgloss.NewStyle().Foreground(lipgloss.Color("9")).Bold(true),
			QuitHint: lipgloss.NewStyle().Foreground(lipgloss.Color("241")).MarginTop(2),
		}
	}
	// Light theme
	return Styles{
		Title:    lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("63")).MarginBottom(1),
		Section:  lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("33")).MarginTop(1).Underline(true),
		Label:    lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("57")),
		Value:    lipgloss.NewStyle().Foreground(lipgloss.Color("235")),
		Error:    lipgloss.NewStyle().Foreground(lipgloss.Color("1")).Bold(true),
		QuitHint: lipgloss.NewStyle().Foreground(lipgloss.Color("240")).MarginTop(2),
	}
}
