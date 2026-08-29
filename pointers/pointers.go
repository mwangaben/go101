package pointers

import "fmt"

func ZeroVal(iVal int) {
	iVal = 0
}

func ZeroPtr(iPtr *int) {
	*iPtr = 0
}

type Person struct {
	Name string
	Age  int
}

func updatePerson(p Person, age int, name string) {
	p.Age = age
	p.Name = name
}

func UpdatePersonPtr(p *Person, age int, name string) {
	p.Age = age
	p.Name = name
}

func ModifyArray(arr *[3]int) {
	(*arr)[0] = 10
	arr[1] = 20
	arr[2] = 30
}

func ModifySlice(s []int) {
	s[0] = 10
	s[1] = 20
	s[2] = 30
}

type Counter struct {
	Count int
}

func (c *Counter) PointerIncrement() *Counter {
	c.Count++
	return c
}

func (c *Counter) PointerDecrement() *Counter {
	c.Count--
	return c
}

type User struct {
	ID    int
	Name  string
	Age   *int
	Email string
}

func UpdateUser(u *User) error {
	if u.Age == nil {
		return fmt.Errorf("the user is nil")
	}
	fmt.Printf("updating the user age to %v \n", *u.Age)
	return nil
}
