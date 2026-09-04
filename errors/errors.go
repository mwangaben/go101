package errors

import (
	"errors"
)

func Add(arg int) (int, error) {
	if arg == 2 {
		return -1, errors.New("can't work with two")
	}
	return arg, nil
}

type Temperature int

const (
	Medium Temperature = iota
	Hot
	VeryHot
)

func MakeTea(cups int, temp Temperature) (int, error) {

	if cups >= 3 {
		return 0, errors.New("you can not make more that two cups at once")
	} else if temp == VeryHot {
		return 0, errors.New("no enough power for now")
	} else {
		return 1, nil
	}
}
