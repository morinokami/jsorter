package jsorter

import (
	"encoding/json"
	"errors"
)

var (
	ErrInvalidJSON = errors.New("Invalid JSON")
)

// Sort sorts the JSON data and returns the result.
func Sort(jsonData []byte, reverse bool) ([]byte, error) {
	if !json.Valid(jsonData) {
		return nil, ErrInvalidJSON
	}

	return []byte(""), nil
}
