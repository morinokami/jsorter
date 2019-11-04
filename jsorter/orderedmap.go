package jsorter

import (
	"bytes"
	"encoding/json"
	"fmt"
)

type item struct {
	key   string
	value interface{}
}

type orderedmap []item

func (om orderedmap) String() string {
	b, _ := om.MarshalJSON()
	return string(b)
}

func (om orderedmap) MarshalJSON() ([]byte, error) {
	var buf bytes.Buffer

	_, err := buf.WriteString("{")
	if err != nil {
		return nil, err
	}

	for i, v := range om {
		if i != 0 {
			// write a delimeter
			_, err = buf.WriteString(",")
			if err != nil {
				return nil, err
			}
		}
		// write a key
		key, err := json.Marshal(v.key)
		if err != nil {
			return nil, err
		}
		_, err = buf.WriteString(fmt.Sprintf(`%s:`, key))
		if err != nil {
			return nil, err
		}
		// write a value
		value, err := json.Marshal(v.value)
		if err != nil {
			return nil, err
		}
		_, err = buf.Write(value)
		if err != nil {
			return nil, err
		}
	}
	if err != nil {
		return nil, err
	}

	_, err = buf.WriteString("}")
	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}
