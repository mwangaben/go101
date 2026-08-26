package maps

import (
	"github.com/go-playground/assert/v2"
	"testing"
)

func TestMaps(t *testing.T) {

	t.Run("it test creation on maps whcih is make(map[key-type]val-type)", func(t *testing.T) {
		m := make(map[string]int)

		m = map[string]int{
			"age":   12,
			"price": 1200,
		}

		assert.Equal(t, m["age"], 12)
		assert.Equal(t, m["price"], 1200)
		assert.Equal(t, 2, len(m))

		delete(m, "age")
		assert.Equal(t, 1, len(m))

		clear(m)
		assert.Equal(t, 0, len(m))
	})

	t.Run("it initialize and assign values", func(t *testing.T) {
		n := map[string]int{"age": 3, "price": 1200}

		assert.Equal(t, 3, n["age"])
		assert.Equal(t, 1200, n["price"])

		value, present := n["height"]
		assert.Equal(t, false, present)
		assert.Equal(t, 0, value)
	})

}
