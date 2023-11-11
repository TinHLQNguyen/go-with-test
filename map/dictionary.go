package main

type Dictionary map[string]string

const (
	ErrNotFound  = DictionaryErr("could not find the word you are looking for")
	ErrWordExist = DictionaryErr("cannot add word because it already exists")
)

type DictionaryErr string

// utilizing error interface. This way these error become immutable
func (e DictionaryErr) Error() string {
	return string(e)
}

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

func (d Dictionary) Update(word, newDefinition string) {
	d[word] = newDefinition
}
