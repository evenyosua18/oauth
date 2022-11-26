package constant

import "errors"

// list of error
var (
	// authentication use case
	ErrInvalidClientSecret = errors.New("invalid client secret")
	ErrInvalidPassword     = errors.New("invalid password")
	ErrInactiveUser        = errors.New("user is not active")
	ErrInvalidScope        = errInvalidScope

	// general
	ErrIP = errors.New("error get ip")
)

func errInvalidScope(scope string) error {
	return errors.New("scope " + scope + " is not found")
}
