package models

import (
	"encoding/json"
	"github.com/quii/go-http-reference-impl/internal/specifications"
)

type PostDTO struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

func NewPostDTOFromPost(post specifications.Post) PostDTO {
	return PostDTO{
		Title:   post.Title,
		Content: post.Content,
	}
}

func (p PostDTO) ToJSON() ([]byte, error) {
	return json.Marshal(p)
}