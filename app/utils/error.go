package utils

import (
	"database/sql"
	"errors"
)

var (
	ErrNoRows                 = sql.ErrNoRows
	ErrNoData                 = errors.New("no data")
	ErrHostResponse           = errors.New("error response")
	ErrExpired                = errors.New("subscription expired")
	ErrUnmatchPassword        = errors.New("unmatch password")
	ErrIncorrectPassword      = errors.New("password is required (min: 6 characters)")
	ErrIncorrectEmail         = errors.New("incorrect email")
	ErrIOAction               = errors.New("I/O action error")
	ErrAccountExists          = errors.New("account already exists")
	ErrAccountAlreadyVerified = errors.New("account is already verified")
	ErrAccountNotFound        = errors.New("account not found")
	ErrInternal               = errors.New("internal error")
	ErrDataBase               = errors.New("data base error")
	ErrNotFound               = errors.New("resource not found")
	ErrNotAllowed             = errors.New("action not allowed")
	ErrForbidden              = errors.New("forbidden access")
	ErrBadConfig              = errors.New("bad config")
	ErrBadToken               = errors.New("bad token")
	ErrBadSession             = errors.New("bad session")
	ErrInvalidation           = errors.New("validation failed")
	ErrUnauthorized           = errors.New("unauthorized")
	ErrBotUnknownCmd          = errors.New("unknown command")
	ErrSendEmail              = errors.New("error send email")
)
