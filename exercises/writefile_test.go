package exercises

import "testing"

func TestOverWrite(t *testing.T) {
	got := overWrite("q.txt", "Hello World")
	if got != nil {
		t.Error("Expected nil Got", got)
	}
}
func TestWriteln(t *testing.T) {
	got := writeln("b.txt", "Hello Globe-")
	if got != nil {
		t.Error("Expected nil Got", got)
	}
}
