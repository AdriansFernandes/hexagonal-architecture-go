package domain

import "errors"

var (
	ErrInvalidName           = errors.New("name must not be empty")
	ErrInvalidCPF            = errors.New("cpf must have 11 digits")
	ErrInvalidAgency         = errors.New("agency must have 4 digits")
	ErrInvalidInitialDeposit = errors.New("initial deposit must not be negative")
)
