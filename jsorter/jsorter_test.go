package jsorter

import (
	"reflect"
	"testing"
)

func TestSort(t *testing.T) {
	ts := []struct {
		txt         string
		orig        []byte
		reverse     bool
		errExpected bool
		exp         interface{}
	}{
		{
			txt:         "invalid JSON",
			orig:        []byte("I'm JSON"),
			errExpected: true,
			exp:         ErrInvalidJSON,
		},
	}

	for _, tc := range ts {
		t.Log(tc.txt)

		result, err := Sort(tc.orig, tc.reverse)

		if tc.errExpected {
			if err == nil {
				t.Error("Expected error, got none.")
			} else if err != tc.exp {
				t.Errorf("Expected %s, got %s", tc.exp, err)
			}
			continue
		}

		if err != nil {
			t.Errorf("Unexpected error: %s", err)
			continue
		}

		if !reflect.DeepEqual(result, tc.exp) {
			t.Error("Not sorted as expected:")
			t.Error(tc.exp)
			t.Error(result)
		}
	}
}
