package pkgs

import "testing"

func TestCompare(t *testing.T) {
	if r := compare([]byte{'2', '3'}, []byte{'2', '3'}); r != 0 {
		t.Error("Expected -1 for compare('2','3') Got ", r)
	}
}
