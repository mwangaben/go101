package goroutines

import "testing"

func TestCallAsyncTask(t *testing.T) {
	t.Run("test the async function ", func(t *testing.T) {

		CallAsyncTask()
	})
}

func TestSelectForRange(t *testing.T) {

	t.Run("It test the select and range to loop over", func(t *testing.T) {
		SelectForRange()
	})
}
