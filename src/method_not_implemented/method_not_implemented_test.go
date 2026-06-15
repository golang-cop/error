package MethodNotImplementedError

import "testing"

func TestNewMessage(t *testing.T) {
	const want = "The method Foo is not implemented"
	e := New("Foo")
	if got := e.Message(); got != want {
		t.Fatalf("Message() = %q, want %q", got, want)
	}
}

func TestNewIsNull(t *testing.T) {
	e := New("Bar")
	if e.IsNull() {
		t.Fatal("IsNull() = true, want false for a not-implemented error")
	}
}

func TestEmptyName(t *testing.T) {
	const want = "The method  is not implemented"
	e := New("")
	if got := e.Message(); got != want {
		t.Fatalf("Message() = %q, want %q", got, want)
	}
}
