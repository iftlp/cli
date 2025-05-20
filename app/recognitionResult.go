package app

import (
	"fmt"
	"main/client"

	tea "github.com/charmbracelet/bubbletea"
)

type apiResultMsg struct {
	result string
	err    error
}

type recognitionResultModel struct {
	filePath string
	client   client.Client

	loading bool
	result  string
	err     error
}

func NewRecognitionResultModel(path string) *recognitionResultModel {
	cl := client.CreateClient("https://shizima-api.baptiiiste.com")
	return &recognitionResultModel{
		filePath: path,
		client:   *cl,
		loading:  true,
	}
}

func sendFileCmd(client client.Client, path string) tea.Cmd {
	return func() tea.Msg {
		res, err := client.SendWavFile(path)
		return apiResultMsg{result: res.Message, err: err}
	}
}

func (m *recognitionResultModel) Init() tea.Cmd {
	return sendFileCmd(m.client, m.filePath)
}

func (m *recognitionResultModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case apiResultMsg:
		m.loading = false
		m.result = msg.result
		m.err = msg.err
		return m, nil
	case tea.KeyMsg:
		if msg.String() == "q" || msg.String() == "ctrl+c" {
			return m, tea.Quit
		}
	}
	return m, nil
}

func (m *recognitionResultModel) View() string {
	if m.loading {
		return "File selected...\n\nWaiting for API response..."
	}
	if m.err != nil {
		return fmt.Sprintf("Error: %v\n\nPress 'q' to quit.", m.err)
	}
	return fmt.Sprintf("API Response:\n\n%s\n\nPress 'q' to quit.", m.result)
}
