package app

import (
	"encoding/json"
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
)

type Artist struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type Album struct {
	ID     int     `json:"id"`
	Title  string  `json:"title"`
	Cover  string  `json:"cover"`
	Artist *Artist `json:"artist"`
}

type Music struct {
	ID            int      `json:"id"`
	Title         string   `json:"title"`
	FilePath      string   `json:"filePath"`
	DurationInSec int      `json:"durationInSec"`
	Artists       []Artist `json:"artists"`
	Album         Album    `json:"album"`
}

type musicInfoModel struct {
	music Music
	err   error
}

func NewMusicInfoModel(jsonString string) *musicInfoModel {
	var music Music
	err := json.Unmarshal([]byte(jsonString), &music)
	if err != nil {
		return &musicInfoModel{err: err}
	}
	return &musicInfoModel{music: music}
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
	if m.err != nil {
		return fmt.Sprintf("Error parsing music info: %v\n\nPress 'q' to quit.", m.err)
	}

	album := m.music.Album
	artistName := "(unknown)"
	if album.Artist != nil {
		artistName = album.Artist.Name
	}

	return fmt.Sprintf(
		"🎵 Music Info\n\n"+
			"Title:        %s\n"+
			"Duration:     %ds\n"+
			"File Path:    %s\n\n"+
			"💿 Album Info\n"+
			"Album:        %s\n"+
			"Cover:        %s\n"+
			"Artist:       %s\n\n"+
			"Press 'q' to quit.",
		m.music.Title,
		m.music.DurationInSec,
		m.music.FilePath,
		album.Title,
		album.Cover,
		artistName,
	)
}
