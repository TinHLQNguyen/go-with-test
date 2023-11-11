package main

import "errors"

type Dictionary map[string]string

var (
	ErrNotFound  = errors.New("could not find the word you are looking for")
	ErrWordExist = errors.New("cannot add word because it already exists")
)

func (d Dictionary) Search(word string) (string, error) {
	definition, ok := d[word]
	if !ok {
		return "", ErrNotFound
	}
	return definition, nil
}

// map type is already pointer, no need here
func (d Dictionary) Add(word, definition string) error {
	_, err := d.Search(word)

	switch err {
	case ErrNotFound:
		d[word] = definition
		return nil
	case nil:
		return ErrWordExist
	default:
		// this should be some unexpected err
		return err
	}
}
