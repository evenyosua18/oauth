package constant

import "errors"

// list of error
var (
	// authentication use case
	ErrInvalidClientSecret = errors.New("invalid client secret")
	ErrInvalidPassword     = errors.New("invalid password")
)
