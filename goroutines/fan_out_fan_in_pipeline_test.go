package goroutines

import "testing"

func TestPipeLineExample(t *testing.T) {

	t.Run("test pipeline", func(t *testing.T) {
		PipeLineExample()
	})

	t.Run("test fan in fan out with error", func(t *testing.T) {
		FanOutWithErrors()
	})
}
