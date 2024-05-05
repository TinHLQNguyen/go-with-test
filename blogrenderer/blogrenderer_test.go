package blogrenderer_test

import (
	"bytes"
	"go-with-test/blogposts"
	"go-with-test/blogrenderer"
	"testing"
)

func TestRender(t *testing.T) {
	var (
		aPost = blogposts.Post{
			Title:       "hello world",
			Description: "This is a description",
			Tags:        []string{"go", "tdd"},
			Body:        "This is a post"}
	)

	t.Run("Convert a single post to HTML", func(t *testing.T) {
		buf := bytes.Buffer{}
		err := blogrenderer.Render(&buf, aPost)

		if err != nil {
			t.Fatal(err)
		}

		got := buf.String()
		want := `<h1>hello world<h1>`
		if got != want {
			t.Errorf("got '%s' want '%s'", got, want)
		}
	})
}
