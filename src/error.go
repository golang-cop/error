package Error

import (
	Object "github.com/go-composites/error/src/object"
	//Result "github.com/go-composites/result/src"
	//String "github.com/go-composites/string/src"
)

type Interface interface {
	Object.Interface
	Message() string
	IsNull() bool
}

type data struct {
	message string
}

func New(message string) Interface {
	return &data{
		message: message,
	}
}

func (d data) Kind() string {
	return Object.Kind(d)
}
func (d data) RespondTo(methodName string) bool {
	return Object.RespondTo(d, methodName)
}
func (d data) Methods() []string {
	return Object.Methods(d)
}
func (d data) Message() string {
	return d.message
}

func (d data) IsNull() bool {
	return false
}
