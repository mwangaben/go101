package pointers

import (
	"encoding/json"
	"fmt"
	"github.com/go-playground/assert/v2"
	assert2 "github.com/stretchr/testify/assert"
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

		data1, _ := json.Marshal(config1)

		ds := `{"host":"localhost","port":8080,"timeout":null}`

		fmt.Printf("the marched config is %v ", string(data1))

		assert.Equal(t, ds, string(data1))

		var actualMap map[string]interface{}
		err := json.Unmarshal(data1, &actualMap)

		assert2.NoError(t, err)

		expectedMap := map[string]interface{}{
			"host":    "localhost",
			"port":    float64(8080),
			"timeout": nil,
		}

		assert.Equal(t, expectedMap, actualMap)
	})

	t.Run("it test the Struct  pointer", func(t *testing.T) {
		age := 34
		benny := &User{ID: 1, Name: "Benedict", Age: &age, Email: "mwangaben@gmail.com"}
		err := UpdateUser(benny)
		assert2.NoError(t, err)

		mylan := &User{
			ID:    2,
			Name:  "Mylan",
			Email: "mylan@uj.com",
		}

		errTwo := UpdateUser(mylan)
		assert2.Error(t, errTwo, "the user is nils")

	})

}
