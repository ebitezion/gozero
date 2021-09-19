package pkgs

import "testing"

/*
*Programmer: techideator@gmail.com

*Purpose:
 Test collections file

*/

func TestGetIndex(t *testing.T) {
	if output := getIndex(1); output != 13 {
		t.Error("Expected 13 from getIndex(1) Got ", output)
	}
}

func TestGetSlice(t *testing.T) {
	//test for n:m
	if output, _ := getSlice(0, 1); output[0] != "Hello" && output[1] != "World" {
		t.Error("Expected hello world from getSlice(1,2) Got", output)
	}

}

func TestAppendToSlice(t *testing.T) {
	expected := []int{0, 0, 2, 3, 5, 7, 0, 0}
	for i, _ := range expected {
		if got := appendToSlice(); got[i] != expected[i] {
			t.Error("Expected", expected, "from appendToSlice()", "Got", got)
		}
	}

}
func TestCopyToSlice(t *testing.T) {
	expected := []int{2, 3, 4, 5, 4, 5}
	result := false
	got := make([]int, 0)
	for i, _ := range expected {

		if got = copyToSlice(); got[i] != expected[i] {
			result = true
		}

	}
	if result {
		t.Error("Expected", expected, "from copyToSlice()", "Got", got)
	}

}

func TestAMap(t *testing.T) {
	expected := map[string]int{
		"Jan": 31, "Feb": 28, "Mar": 31,
		"Apr": 30, "May": 31, "Jun": 30,
		"Jul": 31, "Aug": 31, "Sep": 30,
		"Oct": 31, "Nov": 30, "Dec": 31, // ← Comma required
	}

	result := false
	got := map[string]int{}

	for i, _ := range expected {

		if got = aMap(); got[i] != expected[i] {
			result = true
		}

	}
	if result {
		t.Error("Expected", expected, "from aMap()", "Got", got)
	}

}
