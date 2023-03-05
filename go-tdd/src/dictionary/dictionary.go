package dictionary

import "errors"

type Dictionary map[string]string

func (d Dictionary) Search(word string) (string, error) {
	definition ok := d[worword]

	if !ok {
		return "", errors.New()
	}
	return d[word], nil
}
