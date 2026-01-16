package model

import "time"

// Library represents a Go library to be featured
type Library struct {
	Name        string   `json:"name"`
	Description string   `json:"description"`
	URL         string   `json:"url"`
	Category    string   `json:"category"`
	Tags        []string `json:"tags"`
	Stars       int      `json:"stars,omitempty"`
	Author      string   `json:"author,omitempty"`
}

// PostedLibrary represents a library that has been posted
type PostedLibrary struct {
	Library   Library   `json:"library"`
	PostedAt  time.Time `json:"posted_at"`
	PostID    string    `json:"post_id,omitempty"`
	ImagePath string    `json:"image_path,omitempty"`
}
