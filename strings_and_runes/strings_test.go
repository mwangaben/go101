package strings_and_runes

import (
	"fmt"
	"github.com/go-playground/assert/v2"
	assert2 "github.com/stretchr/testify/assert"
	"strconv"
	"strings"
	"testing"
	"unicode/utf8"
)

func TestString(t *testing.T) {

	t.Run("it study string as the read-only value", func(t *testing.T) {
		s := "Benny"
		assert.Equal(t, 5, len(s))
		assert.Equal(t, 5, utf8.RuneCountInString(s))
	})

	t.Run("it study runes", func(t *testing.T) {
		h := 'H'

		assert2.Equal(t, int32(72), h)
		assert2.Equal(t, "H", string(h))
	})

	t.Run("it concatenation", func(t *testing.T) {
		s3 := "Hello" + " " + "world"
		assert2.Equal(t, "Hello world", s3)
	})

	t.Run("it rune", func(t *testing.T) {
		runes := []rune("Hello")
		fmt.Printf("the slice: %v \n", runes)
		fmt.Printf("the first %v \n", string(runes[0]))
	})
}

func TestStringOperations(t *testing.T) {

	t.Run("it test contains", func(t *testing.T) {
		word := "Hello world"
		assert2.True(t, strings.Contains(word, "Hello"))
	})
	t.Run("it Count letters / runes", func(t *testing.T) {
		word := "Hello world"
		assert2.Equal(t, 3, strings.Count(word, "l"))
	})

	t.Run("it checks Prefix/Suffix", func(t *testing.T) {
		word := "Hello world!"
		assert2.True(t, strings.HasPrefix(word, "Hello"))
		assert2.True(t, strings.HasSuffix(word, "!"))
	})

	t.Run("it check the Index", func(t *testing.T) {
		word := "Hello world!"
		assert2.Equal(t, 6, strings.Index(word, "w"))
		assert2.Equal(t, 11, strings.Index(word, "!"))
	})

	t.Run("It replaces the string", func(t *testing.T) {
		word := "Hello world!"
		assert2.Equal(t, "Hello Go!", strings.Replace(word, "world", "Go", 1))
	})

	t.Run("It split the string into slice", func(t *testing.T) {
		word := "Hello world!"
		expected := []string{"Hello", "world!"}
		assert2.Equal(t, expected, strings.Split(word, " "))
	})

	t.Run("It joins  slice string", func(t *testing.T) {
		slice := []string{"Hello", "world"}
		word := "Hello world"
		assert2.Equal(t, word, strings.Join(slice, " "))
	})

	t.Run("it trim white space", func(t *testing.T) {
		dou := " Hello "
		expected := "Hello"
		assert2.Equal(t, expected, strings.TrimSpace(dou))
	})

	t.Run("It changes cases", func(t *testing.T) {
		word := "Hello"
		expected := "hello"
		upper := "HELLO"
		assert2.Equal(t, expected, strings.ToLower(word))
		assert2.Equal(t, upper, strings.ToUpper(word))

	})

}

func TestConvert(t *testing.T) {

	t.Run("it Convert int to string", func(t *testing.T) {
		n := 42
		assert2.Equal(t, "42", strconv.Itoa(n))
	})

	t.Run("it converts string into int", func(t *testing.T) {
		n := "42"
		atoi, err := strconv.Atoi(n)
		assert2.NoError(t, err)
		assert2.Equal(t, 42, atoi)
	})

	t.Run("it converts Float to string", func(t *testing.T) {
		float := 3.14
		expected := "3.14"
		assert2.Equal(t, expected, strconv.FormatFloat(float, 'f', 2, 64))
	})
}
