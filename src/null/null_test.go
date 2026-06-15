package NullError

import "testing"

func TestNewMessage(t *testing.T) {
	e := New()
	if got := e.Message(); got != "" {
		t.Fatalf("Message() = %q, want empty string", got)
	}
}

func TestNewIsNull(t *testing.T) {
	e := New()
	if !e.IsNull() {
		t.Fatal("IsNull() = false, want true for the null error")
	}
}
