package jsorter

import (
	"encoding/json"
	"reflect"
	"sort"
)

var (
	prefix = ""
	indent = "  "
)

func format(d interface{}) ([]byte, error) {
	return json.MarshalIndent(d, prefix, indent)
}

func sortedKeys(m map[string]interface{}) []string {
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	return keys
}

func sorter(d interface{}) interface{} {
	switch reflect.TypeOf(d).Kind() {
	case reflect.Map:
		// extract keys
		m, _ := d.(map[string]interface{})
		keys := sortedKeys(m)

		// sort the value of each key
		res := orderedmap{}
		for _, key := range keys {
			i := item{key, sorter(m[key])}
			res = append(res, i)
		}

		return res
	case reflect.Slice:
		// sort each items in the slice
		s, _ := d.([]interface{})
		for i, v := range s {
			if v != nil {
				s[i] = sorter(v)
			}
		}

		return s
	default:
		return d
	}
}

// Sort sorts the JSON data and returns the result.
func Sort(d []byte, reverse bool) ([]byte, error) {
	var i interface{}
	if err := json.Unmarshal(d, &i); err != nil {
		return nil, err
	}

	return format(sorter(i))
}
