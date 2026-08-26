package datatypes

import (
	"fmt"
	"github.com/go-playground/assert/v2"
	"testing"
)

func TestFloatNumbers(t *testing.T) {

	t.Run("Check the default of float32", func(t *testing.T) {
		assert.IsEqual(Price, 0)
	})

	t.Run("Check the default of float64", func(t *testing.T) {
		assert.IsEqual(PriceTwo, 0)
	})

	t.Run("Check complex number", func(t *testing.T) {
		assert.IsEqual(ComplexPrice, 0)

		realPart := 23.12222
		imaginaryPart := 43.000002
		ComplexPrice = complex(realPart, imaginaryPart)

		assert.IsEqual(ComplexPrice, realPart+imaginaryPart)
		fmt.Printf("the complex number %v\n", ComplexPrice)

		assert.IsEqual(real(ComplexPrice), realPart)
		assert.IsEqual(imag(ComplexPrice), imaginaryPart)
	})
}
