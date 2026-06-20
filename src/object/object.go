package Object

import (
	"fmt"
	"reflect"
	"strings"
)

type Interface interface {
	/*
		Return the interface kind
	*/
	Kind() string
	/*
		Check if an interface{} respond to a method name
	*/
	RespondTo(string) bool
	Methods() []string
}

type data struct{}

func New() Interface {
	return &data{}
}

func Kind(object interface{}) string {
	data := reflect.TypeOf(object).String()
	kind := strings.Split(data, ".")[0]
	return fmt.Sprintf("%s.Interface", kind)
}

/*
Must be implemented by sub-interface of Object.Interface
*/
func (d data) Kind() string {
	return Kind(d)
}

func RespondTo(object interface{}, methodName string) bool {
	method := reflect.ValueOf(object).MethodByName(methodName)
	//spew.Dump(method)
	return method.IsValid()
}

/*
Must be implemented by sub-interface of Object.Interface
*/
func (d data) RespondTo(methodName string) bool {
	return RespondTo(d, methodName)
}

func Methods(object interface{}) []string {
	t := reflect.TypeOf((object))
	var methods []string
	for i := 0; i < t.NumMethod(); i++ {
		methods = append(methods, t.Method(i).Name)
	}
	return methods
}

func (d data) Methods() []string {
	return Methods(d)
}
