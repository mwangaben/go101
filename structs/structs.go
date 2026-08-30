package structs

type Person struct {
	Name string
	Age  int
}

func NewPerson(name string, age int) *Person {
	p := Person{Name: name, Age: age}
	return &p
}

//Structs with methods

type Rectangle struct {
	Width  int
	Height int
}

func NewRectangle(width, height int) *Rectangle {
	r := Rectangle{Width: width, Height: height}
	return &r
}

func (r *Rectangle) Area() int {
	return r.Width * r.Height
}

func (r *Rectangle) Circumference() int {
	return (r.Width + r.Height) * 2
}
