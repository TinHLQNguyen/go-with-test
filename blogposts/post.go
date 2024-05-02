package blogposts

import (
	"bufio"
	"io"
)

type Post struct {
	Title       string
	Description string
}

func newPost(postFile io.Reader) (Post, error) {
	scanner := bufio.NewScanner(postFile)

	scanner.Scan()
	titleLine := scanner.Text()

	scanner.Scan()
	descriptionLine := scanner.Text()

	err := scanner.Err()
	if err != nil {
		return Post{}, err
	}

	post := Post{Title: string(titleLine)[7:], Description: string(descriptionLine)[13:]} // 7 is for "Title "
	return post, nil
}
