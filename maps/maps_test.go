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
