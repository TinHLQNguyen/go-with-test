package blogrenderer_test

import (
	"bytes"
	"go-with-test/blogposts"
	"go-with-test/blogrenderer"
	"io"
	"testing"

	approvals "github.com/approvals/go-approval-tests"
)

func TestRender(t *testing.T) {
	aPost := blogposts.Post{
		Title:       "hello world",
		Description: "This is a description",
		Tags:        []string{"go", "tdd"},
		Body: `# First recipe!
Welcome to my **amazing recipe blog**. 
I am going to write about stuff`,
	}

	postRenderer, err := blogrenderer.NewPostRenderer()
	if err != nil {
		t.Fatal(err)
	}

	t.Run("Convert a single post to HTML", func(t *testing.T) {
		buf := bytes.Buffer{}

		if err := postRenderer.Render(&buf, aPost); err != nil {
			t.Fatal(err)
		}

		got := buf.String()
		approvals.VerifyString(t, got)
	})
}

func BenchmarkRender(b *testing.B) {
	aPost := blogposts.Post{
		Title:       "hello world",
		Description: "This is a description",
		Tags:        []string{"go", "tdd"},
		Body:        "This is a post",
	}

	postRenderer, err := blogrenderer.NewPostRenderer()
	if err != nil {
		b.Fatal(err)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		postRenderer.Render(io.Discard, aPost)
	}
}

func TestIndex(t *testing.T) {
	t.Run("render index of posts", func(t *testing.T) {
		buf := bytes.Buffer{}
		posts := []blogposts.Post{
			{Title: "Hello World"},
			{Title: "Hello World 2"},
		}

		postRenderer, err := blogrenderer.NewPostRenderer()
		if err != nil {
			t.Fatal(err)
		}

		if err := postRenderer.RenderIndex(&buf, posts); err != nil {
			t.Fatal(err)
		}

		got := buf.String()
		approvals.VerifyString(t, got)
	})
}
