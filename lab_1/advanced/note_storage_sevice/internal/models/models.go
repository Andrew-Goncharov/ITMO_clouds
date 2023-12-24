package models

type Note struct {
	ID      string `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

type NoteData struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}
