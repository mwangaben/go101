package structs

import (
	"encoding/json"
	assert2 "github.com/stretchr/testify/assert"
	"testing"
)

func TestNewPerson(t *testing.T) {
	t.Run("it initial the Person struct", func(t *testing.T) {
		person := NewPerson("Myles", 8)

		assert2.Equal(t, "Myles", person.Name)
		assert2.Equal(t, 8, person.Age)
		person.Age = 12

		assert2.Equal(t, 12, person.Age)
	})

	t.Run("it initial with new ", func(t *testing.T) {
		p := new(Person)
		p.Name = "Mylan"
		p.Age = 3

		assert2.Equal(t, 3, p.Age)
		assert2.Equal(t, "Mylan", p.Name)
	})
}

func TestRectangle(t *testing.T) {

	t.Run("It calculate the area of rect", func(t *testing.T) {
		rect := NewRectangle(12, 10)
		area := rect.Area()
		expectedArea := 120
		circ := rect.Circumference()
		expectedCircumference := 44

		assert2.Equal(t, expectedArea, area)
		assert2.Equal(t, expectedCircumference, circ)
	})

}

func TestComposition(t *testing.T) {

	t.Run("it tests the struct composition", func(t *testing.T) {
		addr := Address{
			Street:  "Kisiwani",
			City:    "DSM",
			Country: "Tanzania",
			ZipCode: "12890",
		}

		emply := Employee{
			Name:    "Mylan Benedict",
			Age:     3,
			Address: addr,
			Email:   "mylanb@uj.com",
		}

		assert2.Equal(t, 3, emply.Age)
		assert2.Equal(t, "Kisiwani", emply.Address.Street)
		assert2.Equal(t, "DSM", emply.Address.City)
	})

}

func TestStructTags(t *testing.T) {

	t.Run("it test struct tags", func(t *testing.T) {
		usr := User{
			ID:       1,
			Name:     "Myrine",
			Email:    "myrineb@uj.com",
			Password: "PasscodeM",
			Age:      1,
		}

		expected := `{"id":1,"name":"Myrine","email":"myrineb@uj.com","age":"1"}`

		marshaled, err := json.Marshal(usr)
		assert2.Equal(t, expected, string(marshaled))
		assert2.NoError(t, err)

		data := &User{}
		err = json.Unmarshal([]byte(expected), data)
		assert2.NoError(t, err)
		assert2.Equal(t, "Myrine", data.Name)
	})
}

func TestCompany(t *testing.T) {

	t.Run("it test Company", func(t *testing.T) {
		company := Company{
			Name: "Uj",
			Address: Address{
				Street:  "Malangalanga",
				City:    "Pwani",
				Country: "Tanzania",
				ZipCode: "23890",
			},
			Employees: []Employee{
				{
					Name: "Benedict",
					Age:  35,
					Address: Address{
						Street:  "Kisiwani",
						City:    "DSM",
						Country: "Tanzania",
						ZipCode: "12890",
					},
					Email: "mwangaben@gmail.com",
				},
			},
		}

		dep := Department{
			Name:          "DevOPS",
			Manager:       "Benedict",
			EmployeeCount: 12,
		}

		assert2.Equal(t, "Benedict", company.Employees[0].Name)
		assert2.Equal(t, "DevOPS", dep.Name)
	})
}
