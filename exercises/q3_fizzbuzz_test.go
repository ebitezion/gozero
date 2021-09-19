package exercises

import "testing"

func TestFizzBuzz(t *testing.T) {
	/*
	 I took a random sample, i expect 3 as a multiple of 3 and it should be printed as fizz;
	 5 as a multiple of 5 and it should be printed as buzz;
	 15 is a multiple of both 5 and 3 and should be printed as fizzbuzz
	*/
	expected := []string{"Fizz", "Buzz", "FizzBuzz"}

	var got []string
	if got = FizzBuzz(); got[2] != expected[1] && got[4] != expected[1] && got[14] != expected[2] {
		t.Error("Expected", expected, "Got", got)
	}
}
