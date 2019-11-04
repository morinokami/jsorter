package jsorter

import (
	"encoding/json"
	"reflect"
	"sort"
	"strings"
)

var (
	prefix = ""
	indent = "  "
)

func format(d interface{}, indent int) ([]byte, error) {
	return json.MarshalIndent(d, prefix, strings.Repeat(" ", indent))
}

func sortedKeys(m map[string]interface{}, reverse bool) []string {
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	if reverse {
		sort.Sort(sort.Reverse(sort.StringSlice(keys)))
	} else {
		sort.Strings(keys)
	}
	return keys
}

func sorter(d interface{}, reverse bool) interface{} {
	switch reflect.TypeOf(d).Kind() {
	case reflect.Map:
		// extract keys
		m, _ := d.(map[string]interface{})
		keys := sortedKeys(m, reverse)

		// sort the value of each key
		res := orderedmap{}
		for _, key := range keys {
			i := item{key, sorter(m[key], reverse)}
			res = append(res, i)
		}

		return res
	case reflect.Slice:
		// sort each items in the slice
		s, _ := d.([]interface{})
		for i, v := range s {
			if v != nil {
				s[i] = sorter(v, reverse)
			}
		}

		return s
	default:
		return d
	}
}

// Sort sorts the JSON data and returns the result.
func Sort(d []byte, reverse bool, indent int) ([]byte, error) {
	var i interface{}
	if err := json.Unmarshal(d, &i); err != nil {
		return nil, err
	}

	return format(sorter(i, reverse), indent)
}
