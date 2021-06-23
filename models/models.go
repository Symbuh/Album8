package models

type Tag struct {
	Name string `json:"name"`
}

type Image struct {
	URL         string `json:"url"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Tags        []Tag
}
