package ports

import "github.com/quii/go-http-reference-impl/internal/specifications"

type Blog interface {
	Publish(post specifications.Post) error
	GetIndex() ([]specifications.IndexPost, error)
}