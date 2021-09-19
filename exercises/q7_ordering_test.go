package exercises

import "testing"

func TestF(t *testing.T) {

	if gotx, goty := f(2, 1); gotx != 1 && goty != 2 {
		t.Error("Expected 1,2 Got", gotx, goty)
	}
	if gotx, goty := f(6, 7); gotx != 6 && goty != 7 {
		t.Error("Expected 6,7 Got", gotx, goty)
	}
	if gotx, goty := f(100, 1); gotx != 1 && goty != 100 {
		t.Error("Expected 1,100 Got", gotx, goty)
	}

}
