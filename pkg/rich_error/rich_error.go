package richerror

import (
	"errors"
)

type OP string

type RichError struct {
	op         OP
	kind       Kind
	message    string
	meta       interface{}
	innerError error
}

func (r *RichError) Error() string {
	return r.message
}

func New(options ...RichErrorOption) *RichError {
	opts := &RichErrorOptions{}
	for _, option := range options {
		option(opts)
	}

	newRichError := &RichError{
		op:         opts.Op,
		kind:       opts.Kind,
		message:    opts.Message,
		meta:       opts.Meta,
		innerError: opts.InnerError,
	}

	return newRichError
}

func (r *RichError) Message() string {
	if r.message != "" {
		return r.message
	}

	if r.innerError == nil {
		return "error message isn't set"
	}

	richErr := new(RichError)
	if errors.As(r.innerError, &richErr) {
		return richErr.Message()
	}

	return r.innerError.Error()
}

func (r *RichError) Kind() Kind {
	if r.kind != "" {
		return r.kind
	}

	richErr := new(RichError)
	if errors.As(r.innerError, &richErr) {
		return richErr.Kind()
	}

	return "kind of error isn't set"
}
