package models

type Tag struct {
	Name string `json:"name"`
}

type Image struct {
	ID          int64    `json:"id"`
	URL         string   `json:"url"`
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Tags        []string `json:"image_tags"`
}
