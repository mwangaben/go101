package errors

import (
	assert2 "github.com/stretchr/testify/assert"
	"testing"
)

func TestUserService_GetUser(t *testing.T) {

	t.Run("it test the Not found error in GetUser", func(t *testing.T) {
		users := NewUserService(map[int]string{
			1: "Myles",
			2: "Mylan",
			3: "Myla",
		})

		user, err := users.GetUser(0)

		assert2.Error(t, err)
		assert2.Equal(t, "", user)
		assert2.ErrorIs(t, err, ErrNotFound)
	})

	t.Run("it tests the invalid input on GetUser", func(t *testing.T) {
		users := NewUserService(map[int]string{
			1: "Neema",
			2: "Merry",
		})

		user, err := users.GetUser(-1)

		assert2.Error(t, err)
		assert2.ErrorIs(t, err, ErrInvalidInput)
		assert2.Equal(t, "", user)
	})

	t.Run("it returns a user when found", func(t *testing.T) {
		users := NewUserService(map[int]string{
			1: "Mylan",
			2: "Myles",
			3: "Myrine",
			4: "Myla",
			5: "Mycah",
		})

		user, err := users.GetUser(2)

		assert2.NoError(t, err)
		assert2.Equal(t, "Myles", user)
	})
}
