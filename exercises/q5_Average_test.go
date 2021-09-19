package exercises

import "testing"

func TestAverage(t *testing.T) {
	expect := 2.00
	if got := Average([]int{1, 2, 3}); got != expect {
		t.Error("Expected", expect, "from Average([]int{1,2,3} Got", got)
	}

}
