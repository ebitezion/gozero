/*
4. Write a program that finds the smallest number
in this list:
x := []int{
 48,96,86,68,
 57,82,63,70,
 37,34,83,27,
 19,97, 9,17,
}
*/package exercises

import (
	"reflect"
	"testing"
)

func TestMin(t *testing.T) {
	//x is the input
	x := []int{
		48, 96, 86, 68,
		57, 82, 63, 70,
		37, 34, 83, 27,
		19, 97, 9, 17,
	}

	//Expected
	expected1 := reflect.Int //for case 1
	expected2 := 9           //for case 2

	//Got
	got, _ := Min(x) //x input sceanrio

	//Test case one
	if reflect.TypeOf(got).Kind() != reflect.Int {
		t.Error("Expected Input slice", x, "\n", ` Expected operation is min(	x := []int{
		48, 96, 86, 68,
		57, 82, 63, 70,
		37, 34, 83, 27,
		19, 97, 9, 17,
	})`, "\n\n", "Expected Output", expected1, "Got", got)
	}
	//Test case two... Testing for the smallest
	if got != expected2 {
		t.Error("Expected Input slice", x, "\n", ` Expected operation is min(	x := []int{
		48, 96, 86, 68,
		57, 82, 63, 70,
		37, 34, 83, 27,
		19, 97, 9, 17,
	}) `, "\n", "Expected Output", expected2, "Got", got)
	}

}
