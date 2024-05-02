package blogposts_test

import (
	"blogposts"
	"errors"
	"io/fs"
	"reflect"
	"testing"
	"testing/fstest"
)

func TestNewBlogPosts(t *testing.T) {
	const (
		firstBody = `Title: Post 1
Description: Description 1
Tags: tdd, go`
		secondBody = `Title: Post 2
Description: Description 2
Tags: rust, borrow-checker`
	)

	fs := fstest.MapFS{
		"hello-word.md":  {Data: []byte(firstBody)},
		"hello-word2.md": {Data: []byte(secondBody)},
	}

	posts, err := blogposts.NewPostsFromFS(fs)

	if err != nil {
		t.Fatal(err)
	}

	got := posts[0]
	want := blogposts.Post{
		Title:       "Post 1",
		Description: "Description 1",
		Tags:        []string{"tdd", "go"},
	}

	assertPost(t, got, want)
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

func assertPost(t *testing.T, got blogposts.Post, want blogposts.Post) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %+v , want %+v ", got, want)
	}
}
