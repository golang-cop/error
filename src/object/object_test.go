package Object_test

import (
	"testing"

	Object "github.com/go-composites/error/src/object"
)

func TestNewKind(t *testing.T) {
	o := Object.New()
	if o == nil {
		t.Fatal("New returned nil")
	}
	if got := o.Kind(); got != "Object.Interface" {
		t.Fatalf("Kind() = %q, want Object.Interface", got)
	}
	if got := Object.Kind(o); got == "" {
		t.Fatal("package Kind() empty")
	}
}

func TestRespondTo(t *testing.T) {
	o := Object.New()
	if !o.RespondTo("Kind") {
		t.Fatal("should respond to Kind")
	}
	if o.RespondTo("NoSuchMethod") {
		t.Fatal("should not respond to NoSuchMethod")
	}
	if !Object.RespondTo(o, "RespondTo") {
		t.Fatal("package RespondTo should find RespondTo")
	}
}

func TestMethods(t *testing.T) {
	o := Object.New()
	if len(o.Methods()) == 0 {
		t.Fatal("Methods() returned none")
	}
	if len(Object.Methods(o)) == 0 {
		t.Fatal("package Methods() returned none")
	}
}
