package jsorter

import (
	"testing"
)

func TestMarshalJSON(t *testing.T) {
	ts := []struct {
		txt  string
		data orderedmap
		exp  string
	}{
		{
			txt:  "Empty map",
			data: orderedmap{},
			exp:  "{}",
		},
		{
			txt: "Map with a key-value pair",
			data: orderedmap{
				{"foo", 1},
			},
			exp: `{"foo":1}`,
		},
		{
			txt: "Map with key-value pairs",
			data: orderedmap{
				{"y", "yyy"},
				{"x", "xxx"},
				{"z", "zzz"},
			},
			exp: `{"y":"yyy","x":"xxx","z":"zzz"}`,
		},
	}

	for _, tc := range ts {
		t.Log(tc.txt)

		result, err := tc.data.MarshalJSON()

		if err != nil {
			t.Errorf("Unexpected error: %s", err)
			continue
		}

		if string(result) != tc.exp {
			t.Error("Not ordered as expected!")
			t.Error("Expected:", tc.exp)
			t.Error("Got:", string(result))
			t.FailNow()
		}
	}
}
