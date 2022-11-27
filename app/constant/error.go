package constant

import "errors"

// string of error
const (
	ErrMessageInvalidExpiredTime = "expired time is invalid"
)

// list of error
var (
	// authentication use case
	ErrInvalidClientSecret = errors.New("invalid client secret")
	ErrInvalidPassword     = errors.New("invalid password")
	ErrInactiveUser        = errors.New("user is not active")
	ErrInvalidScope        = errInvalidScope
	ErrInvalidToken        = errInvalidToken

	// general
	ErrIP = errors.New("error get ip")
)

func errInvalidScope(scope string) error {
	return errors.New("scope " + scope + " is not found")
}

func errInvalidToken(desc string) error {
	return errors.New("invalid token: " + desc)
}
