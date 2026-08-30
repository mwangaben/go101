package structs

// Address Embedded Structs (Composition)
type Address struct {
	Street  string
	City    string
	Country string
	ZipCode string
}
type Employee struct {
	Name    string
	Age     int
	Address Address
	Email   string
}
