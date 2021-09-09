package blog

import (
	"github.com/quii/go-http-reference-impl/internal/specifications"
	"github.com/quii/go-http-reference-impl/internal/testhelpers"
)

type InMemBlog struct {
	posts []specifications.Post
}

func (a *InMemBlog) Publish(post specifications.Post) error {
	a.posts = append(a.posts, post)
	return nil
}

func (a *InMemBlog) GetIndex() ([]specifications.IndexPost, error) {
	var indexPosts []specifications.IndexPost
	for _, posts := range a.posts {
		indexPost := specifications.IndexPost{
			Title: posts.Title,
			ID:    testhelpers.RandomString(),
		}
		indexPosts = append(indexPosts, indexPost)
	}
	return indexPosts, nil
}

