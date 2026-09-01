package enums

type Status int

const (
	Pending Status = iota
	Active
	Inactive
	Deleted
)

func (s Status) String() string {
	return [...]string{"Pending", "Active", "Inactive", "Deleted"}[s]
}
