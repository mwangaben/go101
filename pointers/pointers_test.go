package pointers

import (
	"github.com/go-playground/assert/v2"
	"testing"
)

func TestPointers(t *testing.T) {

	t.Run("It does not change the value", func(t *testing.T) {
		number := 12
		ZeroVal(number)
		assert.Equal(t, 12, number)

		ZeroPtr(&number)
		assert.Equal(t, 0, number)
	})

	t.Run("It does update the original with pointer", func(t *testing.T) {
		pn := Person{Age: 12, Name: "Mylan"}

		updatePerson(pn, 32, "Myles")

		assert.Equal(t, 12, pn.Age)
		assert.Equal(t, "Mylan", pn.Name)

		UpdatePersonPtr(&pn, 42, "Myles")
		assert.Equal(t, 42, pn.Age)
		assert.Equal(t, "Myles", pn.Name)

	})

	t.Run("it update by pointer", func(t *testing.T) {
		arr := [3]int{1, 2, 3}
		ModifyArray(&arr)
		assert.Equal(t, [3]int{10, 20, 30}, arr)

		slice := []int{1, 2, 3}
		ModifySlice(slice)
		assert.Equal(t, []int{10, 20, 30}, slice)
	})

	t.Run("Pointer function", func(t *testing.T) {
		c := Counter{
			Count: 10,
		}
		assert.Equal(t, 10, c.Count)

		c.PointerIncrement()
		assert.Equal(t, 11, c.Count)
		c.PointerDecrement()
		assert.Equal(t, 10, c.Count)

		b := Counter{
			Count: 90,
		}
		b.PointerIncrement().PointerDecrement()
		assert.Equal(t, 90, b.Count)
	})

	t.Run("It has ", func(t *testing.T) {
		users := map[string]*Person{
			"benny": {Name: "Benedict Mwanga", Age: 35},
			"myles": {Name: "Myles Benedict", Age: 8},
			"mylan": {Name: "Mylan Benedict", Age: 3},
		}

		assert.Equal(t, "Benedict Mwanga", users["benny"].Name)
		assert.Equal(t, 3, users["mylan"].Age)
	})

	people := []*Person{
		{Name: "Myrine Benedict", Age: 1},
		{Name: "Myla Benedict", Age: 4},
	}

	assert.Equal(t, "Myrine Benedict", people[0].Name)
	assert.Equal(t, "Myla Benedict", people[1].Name)
	assert.Equal(t, 4, people[1].Age)
}

func TestJson(t *testing.T) {

	t.Run("It test the Json config", func(t *testing.T) {
		port := 8080

		config1 := Config{Host: "localhost", Port: &port}

		assert.Equal(t, port, config1.Port)
	})
}
