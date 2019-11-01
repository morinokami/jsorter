package jsorter

import (
	"encoding/json"
	"errors"
)

var (
	ErrInvalidJSON = errors.New("Invalid JSON")
)

var (
	prefix = ""
	indent = "  "
)

func format(jsonData interface{}) ([]byte, error) {
	return json.MarshalIndent(jsonData, prefix, indent)
}

// Sort sorts the JSON data and returns the result.
func Sort(jsonData []byte, reverse bool) ([]byte, error) {
	if !json.Valid(jsonData) {
		return nil, ErrInvalidJSON
	}

	return []byte(""), nil
}
