package NullError

import (
	Error "github.com/golang-oop/error/src"
)

type Interface interface {
	Error.Interface
}

type data struct {
	message string
}

func New() Interface {
	return &data{
		message: "",
	}
}

func (d data) Message() string {
	return d.message
}

func (d data) IsNull() bool {
	return true
}
