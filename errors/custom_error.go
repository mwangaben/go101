package errors

import "fmt"

type DivisionErr struct {
	Dividend float64
	Divisor  float64
	Message  string
}

func (e DivisionErr) Error() string {
	return fmt.Sprintf("division err: %s (%.2f / %.2f)", e.Message, e.Dividend, e.Divisor)
}

func DivideWithCustomError(a, b float64) (float64, error) {
	if b == 0 {
		return 0, DivisionErr{
			Dividend: a,
			Divisor:  b,
			Message:  "can not divide by zero",
		}
	}
	if a < 0 {
		return 0, DivisionErr{
			Dividend: a,
			Divisor:  b,
			Message:  "dividend can not be negative",
		}
	}

	return a / b, nil
}
