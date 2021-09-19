package pkgs

import (
	"testing"
)

func TestMutatingStr(t *testing.T) {

	if output := mutatingStr("Hello"); output != "Dello" {

		t.Error("Expected Hello, Got", output)
	}

}

func TestReverseString(t *testing.T) {
	expect := "raboof"

	if got := ReverseString("foobar"); got != expect {
		t.Error("Expected ", expect, "from ReverseString(foobar) Got", got)
	}
}
