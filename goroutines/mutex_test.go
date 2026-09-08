package goroutines

import "testing"

func TestAddWithMutex(t *testing.T) {
	t.Run("add with Mutex", func(t *testing.T) {
		AddWithMutex()
	})
}
