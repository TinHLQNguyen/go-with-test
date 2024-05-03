package blogposts

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"strings"
)

const (
	titleSeparator       = "Title: "
	descriptionSeparator = "Description: "
	tagSeparator         = "Tags: "
)

type Post struct {
	Title, Description, Body string
	Tags                     []string
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
	body := readBody(scanner)

	err := scanner.Err()
	if err != nil {
		return Post{}, err
	}

	post := Post{Title: titleLine,
		Description: descriptionLine,
		Tags:        tagsLine,
		Body:        body,
	}
	return post, nil
}

func readBody(scanner *bufio.Scanner) string {
	scanner.Scan() // ignore --- line
	buf := bytes.Buffer{}
	for scanner.Scan() {
		fmt.Fprintln(&buf, scanner.Text()) // this is to write data to buf
	}
	return strings.TrimSuffix(buf.String(), "\n")
}
