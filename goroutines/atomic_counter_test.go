package goroutines

import "testing"

func TestAtomicRace(t *testing.T) {
	t.Run("test simple atomic counter", func(t *testing.T) {

		AtomicRace()
	})
}
