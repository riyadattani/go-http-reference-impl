package blog

import (
	"github.com/quii/go-http-reference-impl/internal/specifications"
	"testing"
)

func TestBlog_Publish(t *testing.T) {
	t.Run("should publish a post", func(t *testing.T) {
		specifications.BlogSpecification(t, &InMemBlog{})
	})
}
