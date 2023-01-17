package oneof

import "errors"

var (
	ErrNoOptionSelected       = errors.New("one option must be selected")
	ErrTooManyOptionsSelected = errors.New("maximum one option can be selected")
)
