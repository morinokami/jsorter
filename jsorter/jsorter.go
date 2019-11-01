package jsorter

import "errors"

var (
	ErrInvalidJSON = errors.New("Invalid JSON")
)

// Sort sorts the JSON data and returns the result.
func Sort(json []byte, reverse bool) ([]byte, error) {
	return []byte(""), nil
}
