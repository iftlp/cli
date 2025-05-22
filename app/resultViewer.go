package app

import (
	"encoding/json"
	"fmt"
	"os"

	"main/model"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type musicInfoModel struct {
	music model.Music
	theme string
	err   error
}

func NewMusicInfoModel(jsonString string) *musicInfoModel {
	var music model.Music
	err := json.Unmarshal([]byte(jsonString), &music)

	theme := os.Getenv("theme")
	if err != nil {
		return &musicInfoModel{err: err, theme: theme}
	}
	return &musicInfoModel{music: music, theme: theme}
}

func (m *musicInfoModel) Init() tea.Cmd {
	return nil
}

func (m *musicInfoModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		if msg.String() == "q" || msg.String() == "ctrl+c" {
			return m, tea.Quit
		}
	}
	return m, nil
}

func (m *musicInfoModel) View() string {
	styles := getStyles(m.theme)

	if m.err != nil {
		return styles.Error.Render(fmt.Sprintf("Error parsing music info: %v\n\nPress 'q' to quit.", m.err))
	}

	album := m.music.Album
	artistName := "(unknown)"
	if album.Artist != nil {
		artistName = album.Artist.Name
	}

	return lipgloss.JoinVertical(lipgloss.Left,
		styles.Title.Render("🎵 Music Info"),
		fmt.Sprintf("%s %s", styles.Label.Render("Title:"), styles.Value.Render(m.music.Title)),
		fmt.Sprintf("%s %s", styles.Label.Render("Duration:"), styles.Value.Render(fmt.Sprintf("%ds", m.music.DurationInSec))),
		fmt.Sprintf("%s %s", styles.Label.Render("File Path:"), styles.Value.Render(m.music.FilePath)),

		styles.Section.Render("💿 Album Info"),
		fmt.Sprintf("%s %s", styles.Label.Render("Album:"), styles.Value.Render(album.Title)),
		fmt.Sprintf("%s %s", styles.Label.Render("Cover:"), styles.Value.Render(album.Cover)),
		fmt.Sprintf("%s %s", styles.Label.Render("Artist:"), styles.Value.Render(artistName)),

		styles.QuitHint.Render("Press 'q' to quit."),
	)
}
