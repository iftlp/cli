package app

import (
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
)

type menuItem struct {
	name        string
	description string
}

func (m menuItem) Title() string       { return m.name }
func (m menuItem) Description() string { return m.description }
func (m menuItem) FilterValue() string { return m.name }

type menuModel struct {
	list list.Model
}

func NewMenuModel() menuModel {
	delegate := list.NewDefaultDelegate()
	var items []list.Item

	items = append(items, menuItem{
		name:        "Use wav file",
		description: "Research a song using a pre recorded audio",
	})
	items = append(items, menuItem{
		name:        "Record from mic",
		description: "Research a song recorded from the microphone",
	})

	l := list.New(items, delegate, 60, 20)
	l.Title = "Shizima - Menu"

	menu := menuModel{
		list: l,
	}

	return menu
}

func (m menuModel) Init() tea.Cmd {
	return nil
}

func (m menuModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "enter":
			selected := m.list.SelectedItem().(menuItem)
			if selected.name == "Use wav file" {
				return NewBrowserModel(), nil
			}
		case "q", "ctrl+c":
			return m, tea.Quit
		}
	}

	m.list, cmd = m.list.Update(msg)
	return m, cmd
}

func (m menuModel) View() string {
	return m.list.View()
}
