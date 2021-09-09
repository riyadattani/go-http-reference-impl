// +build acceptance

package acceptance

import (
	"github.com/quii/go-http-reference-impl/internal/specifications"
	"github.com/quii/go-http-reference-impl/internal/specifications/adapters"
	"testing"
)

func TestBlog(t *testing.T) {
	t.Run("When an author publishes a post, I can view it", func(t *testing.T) {
		client := adapters.NewAPIClient(getBaseURL(t), t)
		specifications.BlogSpecification(t, client)
	})
}