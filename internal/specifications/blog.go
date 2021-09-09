package specifications

import (
	"github.com/matryer/is"
	"github.com/quii/go-http-reference-impl/internal/testhelpers"
	"testing"
)

type Post struct {
	Title   string
	Content string
}

type IndexPost struct {
	Title string
	ID    string
}

type BlogAdapter interface {
	Publish(post Post) error
	GetIndex() ([]IndexPost, error)
}

func BlogSpecification(t *testing.T, adapter BlogAdapter) {
	t.Helper()
	t.Run("when an author publishes a post, it appears in the index", func(t *testing.T) {
		is := is.New(t)
		//Given
		//a post
		post := Post{
			Title:   testhelpers.RandomString(),
			Content: testhelpers.RandomString(),
		}

		//When
		// I publish it
		err := adapter.Publish(post)
		is.NoErr(err)

		//Then
		index, err := adapter.GetIndex()
		is.NoErr(err)

		assertPostIsInIndex(t, index, post)
	})
}

func assertPostIsInIndex(t *testing.T, index []IndexPost, post Post) {
	t.Helper()
	found := false
	for _, indexPost := range index {
		if indexPost.Title == post.Title {
			found = true
			break
		}
	}

	if !found {
		t.Fatalf("Could not find post %s", post.Title)
	}
}
