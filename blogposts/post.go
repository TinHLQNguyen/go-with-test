package blogposts

import (
	"bufio"
	"io"
)

const (
	titleSeparator       = "Title: "
	descriptionSeparator = "Description: "
)

type Post struct {
	Title       string
	Description string
}

func newPost(postFile io.Reader) (Post, error) {
	scanner := bufio.NewScanner(postFile)

	readLine := func() string {
		scanner.Scan()
		return scanner.Text()
	}

	titleLine := readLine()[len(titleSeparator):]
	descriptionLine := readLine()[len(descriptionSeparator):]

	err := scanner.Err()
	if err != nil {
		return Post{}, err
	}

	post := Post{Title: titleLine, Description: descriptionLine} // 7 is for "Title "
	return post, nil
}
