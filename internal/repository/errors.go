package repository

import "errors"

var (
	ErrClientIsNil      = errors.New("client is nil")
	ErrEmptyArgument    = errors.New("empty argument")
	ErrUnexpectedResult = errors.New("unexpected result")
)
