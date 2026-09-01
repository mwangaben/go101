package pointers

type WeekEnd int

const (
	Saturday WeekEnd = iota
	Sunday
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
)

func (d WeekEnd) String() string {
	names := []string{"Saturday", "Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "friday"}

	if int(d) < 0 || int(d) >= len(names) {
		return "Unknown"
	}
	return names[d]
}

func (d WeekEnd) IsWeekEnd() bool {
	return d == Sunday || d == Saturday
}

func (d WeekEnd) IsWeekDay() bool {
	return !d.IsWeekEnd()
}
