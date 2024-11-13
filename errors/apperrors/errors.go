package apperrors

import "errors"

var ErrInvalidEnv = errors.New("invalid env")

var ErrUrlNotFound = errors.New("url not found")
