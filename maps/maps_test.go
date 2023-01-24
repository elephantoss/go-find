package maps_test

import (
	"testing"
	"testing/quick"

	"github.com/elephantoss/go-find/maps"
)

func TestFind(t *testing.T) {
	config := &quick.Config{
		MaxCount: 100,
	}

	f := func(c map[string]string, k string, d string) bool {
		v := maps.Find(c, k, d)

		e, ok := c[k]

		return (ok && v == e) || (!ok && v == d)
	}

	if err := quick.Check(f, config); err != nil {
		t.Error(err)
	}
}

func TestFind_Table(t *testing.T) {
	cases := []struct {
		title         string
		input_map     map[string]int
		key           string
		default_value int
		expectation   int
	}{
		{
			title:         "return default if key is not in the map",
			input_map:     map[string]int{"a": 10},
			key:           "b",
			default_value: 50,
			expectation:   50,
		},
		{
			title:         "return value from map if key is in the map",
			input_map:     map[string]int{"a": 10},
			key:           "a",
			default_value: 50,
			expectation:   10,
		},
	}

	for _, tc := range cases {
		t.Run(tc.title, func(t *testing.T) {
			actual := maps.Find(tc.input_map, tc.key, tc.default_value)
			if actual != tc.expectation {
				t.Fatalf(
					"%s: expected: %d got: %d",
					tc.title, tc.expectation, actual)
			}
		})
	}
}
