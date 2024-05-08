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
	var (
		aPost = blogposts.Post{
			Title:       "hello world",
			Description: "This is a description",
			Tags:        []string{"go", "tdd"},
			Body:        "This is a post"}
	)

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
	var (
		aPost = blogposts.Post{
			Title:       "hello world",
			Description: "This is a description",
			Tags:        []string{"go", "tdd"},
			Body:        "This is a post"}
	)

	postRenderer, err := blogrenderer.NewPostRenderer()
	if err != nil {
		b.Fatal(err)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		postRenderer.Render(io.Discard, aPost)
	}
}
