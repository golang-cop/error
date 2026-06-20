package MethodNotImplementedError

import (
	"fmt"

	Error "github.com/go-composites/error/src"
	Object "github.com/go-composites/error/src/object"
)

type Interface interface {
	Error.Interface
}

type data struct {
	message string
}

func New(name string) Interface {
	message := fmt.Sprintf(
		"The method %s is not implemented",
		name,
	)
	return &data{
		message: message,
	}
}

func (d data) Kind() string {
	return Object.Kind(d)
	//return `MethodNotImplementedError.Interface`
}
func (d data) RespondTo(name string) bool {
	return Object.RespondTo(d, name)
}

func (d data) Message() string {
	return d.message
}

func (d data) IsNull() bool {
	return false
}

func (d data) Methods() []string {
	return Object.Methods(d)
}
