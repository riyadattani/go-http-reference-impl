package blog_handler

import (
	"encoding/json"
	"github.com/quii/go-http-reference-impl/internal/ports"
	"github.com/quii/go-http-reference-impl/internal/specifications"
	"net/http"
)

type BlogHandler struct {
	blog ports.Blog
}

func NewBlogHandler(blog ports.Blog) *BlogHandler {
	return &BlogHandler{blog: blog}
}

func (b *BlogHandler) Publish(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func (b *BlogHandler) GetIndex (w http.ResponseWriter, r *http.Request) {
	indexPost := specifications.IndexPost{
		Title: "blah",
		ID:    "blah",
	}
	indexPosts := []specifications.IndexPost{indexPost}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(indexPosts)
}


