package NullError

import (
	Error "github.com/golang-cop/error/src"
	Object "github.com/golang-cop/error/src/object"
)

type Interface interface {
	Object.Interface
	Error.Interface
}

type data struct{}

func New() Interface {
	return &data{}
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
