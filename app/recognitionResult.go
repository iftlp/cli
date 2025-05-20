package app

import (
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

	result string
	err    error
}

func NewRecognitionResultModel(path string) *recognitionResultModel {
	cl := client.CreateClient("https://shizima-api.baptiiiste.com")
	return &recognitionResultModel{
		filePath: path,
		client:   *cl,
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
		m.result = msg.result
		m.err = msg.err
		return NewMusicInfoModel(m.result), nil
	}
	return m, nil
}

func (m *recognitionResultModel) View() string {
	return "File selected...\n\nWaiting for API response..."
}
