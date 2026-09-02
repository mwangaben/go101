package structs

import (
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestAnimal(t *testing.T) {
	t.Run("Animal", func(t *testing.T) {
		a := NewAnimal("Animalia")

		d := NewDog("Dog", *a)
		assert.Equal(t, a.Name, "Animalia")
		assert.Equal(t, "make sound", a.Speak())

		assert.Equal(t, "Baking", d.Speak())
		assert.Equal(t, "Animalia", d.Animal.Name)
	})
}
