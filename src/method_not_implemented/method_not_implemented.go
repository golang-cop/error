package MethodNotImplementedError

import (
	"fmt"

	Error "github.com/golang-oop/error/src"
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

func (d data) Message() string {
	return d.message
}

func (d data) IsNull() bool {
	return false
}
