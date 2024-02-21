package richerror

type Kind int

const (
	ErrKindNotFound Kind = iota + 1
	ErrKindInvalid
	ErrKindForbidden
	ErrKindUnExpected
	ErrKindBadRequest
)
