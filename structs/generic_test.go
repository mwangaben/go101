package structs

import (
	assert2 "github.com/stretchr/testify/assert"
	"testing"
)

func TestGenericStruct(t *testing.T) {

	t.Run("test", func(t *testing.T) {
		intPair := Pair[int]{
			First:  10,
			Second: 12,
		}

		stringPair := Pair[string]{
			First:  "Benedict",
			Second: "Mwanga",
		}

		assert2.Equal(t, 10, intPair.First)
		assert2.Equal(t, 12, intPair.Second)

		assert2.Equal(t, "Benedict", stringPair.First)
		assert2.Equal(t, "Mwanga", stringPair.Second)
	})
}
