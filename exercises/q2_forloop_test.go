package exercises

import "testing"

//Test the q2_forloop exercise
func TestLoopFillsArrayAndPrint(t *testing.T) {
	expected := [10]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	var got [10]int
	result := true
	for i := 0; i < 10; i++ {
		if got = LoopFillsArrayAndPrint(); got[i] != expected[i] {
			result = false
		}
	}
	if !result {
		t.Error("Expected", expected, "from LoopFillsArrayAndPrint()", "Got", got)

	}
}

func TestForLoopTenTimes(t *testing.T) {
	expected := 10
	got := 0
	result := true
	for i := 0; i < 10; i++ {
		if got = forLoopTenTimes(); got != expected {
			result = false
		}
	}
	if !result {
		t.Error("Expected", expected, "from forLoopTenTimes()", "Got", got)

	}
}

func TestGotoLoopTenTimes(t *testing.T) {
	expected := 10
	got := 0
	result := true
	for i := 0; i < 10; i++ {
		if got = gotoLoopTenTimes(); got != expected {
			result = false
		}
	}
	if !result {
		t.Error("Expected", expected, "from gotoLoopTenTimes()", "Got", got)

	}
}
