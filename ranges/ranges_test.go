package ranges

import (
	"github.com/go-playground/assert/v2"
	"testing"
)

func TestRanges(t *testing.T) {

	t.Run("It test range sum", func(t *testing.T) {
		numbers := []int{12, 12, 12}

		assert.Equal(t, 36, Sums(numbers))
	})

	t.Run("It show values and keys from the maps", func(t *testing.T) {
		info := map[string]string{
			"name":     "Benedict",
			"location": "Tanzania",
			"job":      "Developer",
		}

		props, values := KVs(info)
		assert.Equal(t, []string{"name", "location", "job"}, props)
		assert.Equal(t, []string{"Benedict", "Tanzania", "Developer"}, values)
	})
}
