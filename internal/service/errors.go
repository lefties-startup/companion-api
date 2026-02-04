package service

import "github.com/pkg/errors"

var (
	NilLoggerErr     = errors.New("logger is nil")
	NilServiceErr    = errors.New("service is nil")
	NilRepositoryErr = errors.New("repository is nil")
	NotFoundErr      = errors.New("not found")
)
