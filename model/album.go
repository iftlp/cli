package model

type Album struct {
	ID     int     `json:"id"`
	Title  string  `json:"title"`
	Cover  string  `json:"cover"`
	Artist *Artist `json:"artist"`
}
