package goroutines

import "testing"

func TestProcessJob(t *testing.T) {

	t.Run("testing", func(t *testing.T) {
		ProcessJob(10)
	})
}
