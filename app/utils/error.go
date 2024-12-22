package utils

import (
	"database/sql"
	"errors"
)

var (
	ErrNoRows         = sql.ErrNoRows
	ErrNoData         = errors.New("no data")
	ErrIncorrectLogin = errors.New("The login is invalid")
	ErrIncorrectEmail = errors.New("The email address is invalid")
	ErrInternal       = errors.New("internal error")
	ErrDataBase       = errors.New("data base error")
	ErrBadConfig      = errors.New("bad config")
	ErrBadToken       = errors.New("bad token")
	ErrBadSession     = errors.New("bad session")
	ErrSendEmail      = errors.New("error send email")
)
