package main

import (
	"fmt"

	Error "github.com/golang-cop/error/src"
	MethodNotImplementedError "github.com/golang-cop/error/src/method_not_implemented"
	NullError "github.com/golang-cop/error/src/null"
	Object "github.com/golang-cop/error/src/object"
)

func main() {
	o := Object.New()
	e := Error.New(`something gone wrong`)
	n := NullError.New()
	mni := MethodNotImplementedError.New(`boo`)
	fmt.Printf("e.RespondTo(`toto`) %t\n", e.RespondTo(`toto`))
	fmt.Printf("e.RespondTo(`Kind`) %t\n", e.RespondTo(`Kind`))
	fmt.Printf("o.Kind(): %s\n", o.Kind())
	fmt.Printf("e.Kind(): %s\n", e.Kind())
	fmt.Printf("n.Kind(): %s\n", n.Kind())
	fmt.Printf("o.Methods(): %+v \n", o.Methods())
	fmt.Printf("e.Methods(): %+v \n", e.Methods())
	fmt.Printf("mni.Methods(): %+v \n", mni.Methods())
}
