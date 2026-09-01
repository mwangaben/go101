package enums

import "fmt"

type HttpStatus int

const (
	OK                  HttpStatus = 200
	Created             HttpStatus = 201
	Accepted            HttpStatus = 202
	BadRequest          HttpStatus = 400
	Unauthorized        HttpStatus = 401
	Forbidden           HttpStatus = 403
	NotFound            HttpStatus = 404
	InternalServerError HttpStatus = 500
)

func (s HttpStatus) String() string {
	switch s {
	case OK:
		return "200 Ok"
	case Created:
		return "201 Created"
	case Accepted:
		return "202 Accepted"
	case BadRequest:
		return "400 Bad request"
	case Unauthorized:
		return "401 This action is Authorized Action"
	case Forbidden:
		return "403 Forbidden"
	case NotFound:
		return "404 Not Found"
	case InternalServerError:
		return "500 Internal Server error"
	default:
		return fmt.Sprintf("%d Unknown", s)

	}
}

func (s HttpStatus) IsSuccess() bool {
	return s >= 200 && s < 300
}

func (s HttpStatus) IsClientError() bool {
	return s >= 400 && s < 500
}
func (s HttpStatus) IsServerError() bool {
	return s >= 500 && s < 600
}
