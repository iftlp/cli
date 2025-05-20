package app

import (
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
)

type fileItem struct {
	name  string
	path  string
	isDir bool
}

func (f fileItem) Title() string       { return f.name }
func (f fileItem) Description() string { return f.path }
func (f fileItem) FilterValue() string { return f.name }

type browserModel struct {
	list        list.Model
	currentPath string
}

func NewBrowserModel() browserModel {
	startPath, _ := os.Getwd()
	items := readDir(startPath)

	delegate := list.NewDefaultDelegate()
	l := list.New(items, delegate, 60, 20)
	l.Title = "Select a file"

	return browserModel{
		list:        l,
		currentPath: startPath,
	}
}

func readDir(path string) []list.Item {
	entries, err := ioutil.ReadDir(path)
	if err != nil {
		return []list.Item{}
	}

	var items []list.Item

	if filepath.Dir(path) != path {
		items = append(items, fileItem{
			name:  "..",
			path:  filepath.Dir(path),
			isDir: true,
		})
	}

	for _, entry := range entries {
		items = append(items, fileItem{
			name:  entry.Name(),
			path:  filepath.Join(path, entry.Name()),
			isDir: entry.IsDir(),
		})
	}
	return items
}

func (m browserModel) Init() tea.Cmd {
	return nil
}

func (m browserModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "enter":
			selected := m.list.SelectedItem().(fileItem)
			if selected.isDir {
				m.currentPath = selected.path
				m.list.SetItems(readDir(m.currentPath))
			} else {
				m.list.Title = "Song recognition"
				newModel := NewRecognitionResultModel(selected.path)
				return newModel, newModel.Init()
			}
		case "q", "ctrl+c":
			return m, tea.Quit
		}
	}

	m.list, cmd = m.list.Update(msg)
	return m, cmd
}

func (m browserModel) View() string {
	return m.list.View()
}
