package structs

type Animal struct {
	Name string
}

func (a Animal) Eat() string {
	return "Eating.."
}

func (a Animal) Move() string {
	return "working"
}

func (a Animal) Speak() string {
	return "make sound"
}

func NewAnimal(name string) *Animal {
	return &Animal{name}
}

type Dog struct {
	Animal   Animal
	Breeding string
}

func (d Dog) Speak() string {
	return "Baking"
}

func (d Dog) Move() string {
	return "working"
}

func (d Dog) Eat() string {
	return "eating"
}

func NewDog(name string, animal Animal) *Dog {
	return &Dog{Breeding: name, Animal: animal}
}
