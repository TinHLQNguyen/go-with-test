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
		"hello-word.md":  {Data: []byte("Title: Post 1")},
		"hello-word2.md": {Data: []byte("Title: Post 2")},
	}

	posts, err := blogposts.NewPostsFromFS(fs)

	if err != nil {
		t.Fatal(err)
	}

	got := posts[0]
	want := blogposts.Post{Title: "Post 1"}

	if got != want {
		t.Errorf("got %+v posts, want %+v posts", got, want)
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
