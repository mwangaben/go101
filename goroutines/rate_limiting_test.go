package goroutines

import "testing"

func TestRateLimiting(t *testing.T) {
	t.Run("testing simple rate limiting via channel and goroutines", func(t *testing.T) {
		SimpleRateLimiter()
	})

	t.Run("testing burst rate limiter via channel and goroutines", func(t *testing.T) {
		BurstRateLimiter()
	})
}
