package link

import "errors"

var (
	ErrCodeTaken    = errors.New("code already taken")
	ErrInvalidInput = errors.New("invalid input")
	ErrNotFound     = errors.New("link not found")
	ErrExpired      = errors.New("link expired")
)
