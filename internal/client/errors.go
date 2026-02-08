package client

import (
	"errors"
)

var (
	ErrEmptyClient     = errors.New("client is nil")
	ErrConnectionIsNil = errors.New("connection is nil")
	ErrArgumentIsNil   = errors.New("argument is nil")
)
