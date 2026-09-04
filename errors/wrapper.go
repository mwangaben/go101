package errors

import (
	"errors"
	"fmt"
)

type DatabaseErr struct {
	Err     error
	Query   string
	Message string
}

func (d DatabaseErr) Error() string {
	return fmt.Sprintf("database err: %v  (query: %v)", d.Message, d.Query)
}

func (e DatabaseErr) UnWrap() error {
	return e.Err
}

func QueryDatabase(query string) error {
	return DatabaseErr{
		Err:     errors.New("connection timeout"),
		Query:   query,
		Message: "Failed to execute query",
	}
}
func ProcessQuery(query string) error {
	err := QueryDatabase(query)

	if err != nil {
		return fmt.Errorf("processing query %s: %w ", query, err)
	}

	return nil
}
