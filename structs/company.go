package structs

type Company struct {
	Name      string
	Address   Address
	Employees []Employee
}
type Department struct {
	Name          string
	Manager       string
	EmployeeCount int
}
