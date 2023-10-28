package models

type Project struct {
	ID          string  `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Authors     string  `json:"authors"`
	Images      []Image `json:"images"`
}
