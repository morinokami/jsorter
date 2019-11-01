package jsorter

import (
	"testing"
)

func TestSort(t *testing.T) {
	ts := []struct {
		txt         string
		orig        string
		reverse     bool
		errExpected bool
		exp         interface{}
	}{
		{
			txt:         "invalid JSON 1",
			orig:        "I'm JSON",
			errExpected: true,
			exp:         ErrInvalidJSON,
		},
		{
			txt:         "invalid JSON 2",
			orig:        `{"example":2:]}}`,
			errExpected: true,
			exp:         ErrInvalidJSON,
		},
		// https://json.org/example.html
		{
			txt: "valid JSON 1",
			orig: `{
  "glossary":{
    "title":"example glossary",
    "GlossDiv":{
      "title":"S",
      "GlossList":{
        "GlossEntry":{
          "ID":"SGML",
          "SortAs":"SGML",
          "GlossTerm":"Standard Generalized Markup Language",
          "Acronym":"SGML",
          "Abbrev":"ISO 8879:1986",
          "GlossDef":{
            "para":"A meta-markup language, used to create markup languages such as DocBook.",
            "GlossSeeAlso":[
              "GML",
              "XML"
            ]
          },
          "GlossSee":"markup"
        }
      }
    }
  }
}`,
			exp: `{
  "glossary":{
    "GlossDiv":{
      "GlossList":{
        "GlossEntry":{
          "Abbrev":"ISO 8879:1986",
          "Acronym":"SGML",
          "GlossDef":{
            "GlossSeeAlso":[
              "GML",
              "XML"
            ],
            "para":"A meta-markup language, used to create markup languages such as DocBook."
          },
          "GlossSee":"markup",
          "GlossTerm":"Standard Generalized Markup Language",
          "ID":"SGML",
          "SortAs":"SGML"
        }
      },
      "title":"S"
    },
    "title":"example glossary"
  }
}`,
		},
	}

	for _, tc := range ts {
		t.Log(tc.txt)

		result, err := Sort([]byte(tc.orig), tc.reverse)

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

		if string(result) != tc.exp {
			t.Error("Not sorted as expected:")
			t.Error("Expected:", tc.exp)
			t.Error("Got:", result)
		}
	}
}
