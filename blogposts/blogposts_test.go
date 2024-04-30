package blogposts_test

import (
	"blogposts"
	"errors"
	"io/fs"
	"testing"
	"testing/fstest"
)

func TestNewBlogPosts(t *testing.T) {
	fs := fstest.MapFS{
		"hello-word.md":  {Data: []byte("hi")},
		"hello-word2.md": {Data: []byte("hola")},
	}

	posts, err := blogposts.NewPostsFromFS(fs)

	if err != nil {
		t.Fatal(err)
	}

	got := len(posts)
	want := len(fs)

	if got != want {
		t.Errorf("got %d posts, want %d posts", got, want)
	}
}

func TestOpenErrorFS(t *testing.T) {
	_, err := blogposts.NewPostsFromFS(StubFailingFS{})

	if err == nil {
		t.Errorf("This should be an error")
	}
}

type StubFailingFS struct {
}

func (s StubFailingFS) Open(name string) (fs.File, error) {
	return nil, errors.New("oh no, i always fail")
}
