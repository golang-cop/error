package NullError

import (
	Error "github.com/go-composites/error/src"
	Object "github.com/go-composites/error/src/object"
)

type Interface interface {
	Object.Interface
	Error.Interface
}

type data struct{}

// theNullError is the single, immutable "no error" sentinel — empty message,
// IsNull() always true, so every NullError is interchangeable: one shared
// instance is semantically identical to a fresh one, at zero alloc.
var theNullError = &data{}

func New() Interface {
	return theNullError
}

func (d data) Kind() string {
	//return `NullError.Interface`
	return Object.Kind(d)
}
func (d data) RespondTo(methodName string) bool {
	return Object.RespondTo(d, methodName)
}

func (d data) Message() string {
	return ``
}

func (d data) IsNull() bool {
	return true
}

func (d data) Methods() []string {
	return Object.Methods(d)
}
