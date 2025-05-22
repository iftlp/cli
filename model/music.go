package model

type Music struct {
	ID            int      `json:"id"`
	Title         string   `json:"title"`
	FilePath      string   `json:"filePath"`
	DurationInSec int      `json:"durationInSec"`
	Artists       []Artist `json:"artists"`
	Album         Album    `json:"album"`
}
