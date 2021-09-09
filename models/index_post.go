package models

import (
	"encoding/json"
	"github.com/quii/go-http-reference-impl/internal/specifications"
)

type IndexPostDTO struct {
	Title string `json:"title"`
	ID    string `json:"id"`
}

func NewIndexPostListFromJSON(in []byte) ([]specifications.IndexPost, error) {
	var indexPostDTOs []IndexPostDTO

	err := json.Unmarshal(in, &indexPostDTOs)
	if err != nil {
		return nil, err
	}

	var indexPosts []specifications.IndexPost

	for _, dto := range indexPostDTOs {
		indexPost := specifications.IndexPost{
			Title: dto.Title,
			ID:    dto.Title,
		}
		indexPosts = append(indexPosts, indexPost)
	}
	return indexPosts, nil
}
