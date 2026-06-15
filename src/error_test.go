package Error

import "testing"

func TestNewMessage(t *testing.T) {
	const want = "boom"
	e := New(want)
	if got := e.Message(); got != want {
		t.Fatalf("Message() = %q, want %q", got, want)
	}
}

func TestNewIsNull(t *testing.T) {
	e := New("anything")
	if e.IsNull() {
		t.Fatal("IsNull() = true, want false for a real error")
	}
}

func TestEmptyMessage(t *testing.T) {
	e := New("")
	if got := e.Message(); got != "" {
		t.Fatalf("Message() = %q, want empty string", got)
	}
	if e.IsNull() {
		t.Fatal("IsNull() = true, want false even for empty message")
	}
}
