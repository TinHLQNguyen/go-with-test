package blogposts

import (
	"bufio"
	"io"
	"strings"
)

const (
	titleSeparator       = "Title: "
	descriptionSeparator = "Description: "
	tagSeparator         = "Tags: "
)

type Post struct {
	Title       string
	Description string
	Tags        []string
}

func newPost(postFile io.Reader) (Post, error) {
	scanner := bufio.NewScanner(postFile)

	readMetaLine := func(tagName string) string {
		scanner.Scan()
		return strings.TrimPrefix(scanner.Text(), tagName)
	}

	titleLine := readMetaLine(titleSeparator)
	descriptionLine := readMetaLine(descriptionSeparator)
	tagsLine := strings.Split(readMetaLine(tagSeparator), ", ")

	err := scanner.Err()
	if err != nil {
		return Post{}, err
	}

	post := Post{Title: titleLine, Description: descriptionLine, Tags: tagsLine}
	return post, nil
}
