package models

type Image struct {
	ID        string `json:"id"`
	CreatedAt string `json:"created_at"`
	Filename  string `json:"path"`
}
