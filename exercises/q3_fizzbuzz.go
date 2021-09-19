package exercises

import "strconv"

/*
Q3. (0) FizzBuzz
1. Solve this problem, called the Fizz-Buzz [23] problem:
Write a program that prints the numbers from 1 to 100. But for multiples
of three print “Fizz” instead of the number and for the multiples of five
print “Buzz”. For numbers which are multiples of both three and five print
“FizzBuzz”.
*/

func FizzBuzz() []string {
	var record []string
	for i := 1; i < 101; i++ {

		if i%3 == 0 && i%5 == 0 {
			record = append(record, "Fizz Buzz")
		} else if i%3 == 0 {
			//multiple of 3
			record = append(record, "Fizz")
		} else if i%5 == 0 {
			//multiples of five
			record = append(record, "Buzz")
		} else {
			record = append(record, strconv.Itoa(i))
		}

	}
	return record
}
