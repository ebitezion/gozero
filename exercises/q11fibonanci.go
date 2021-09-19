/*
Package exercises ...
Q11. (1) Fibonacci
1. The Fibonacci sequence starts as follows: 1, 1, 2, 3, 5, 8, 13, . . . Or in mathematical
terms: x1 = 1; x2 = 1; xn = xn−1 + xn−2 ∀n > 2.
Write a function that takes an int value and gives that many terms of the Fibonacci
sequence*/
package exercises

/*compute financi of the order x(n)=x(n)-1 +x(n)-2...*/
func Fib(n int) int {
	/*using recursion*/
	//terminating condition is Fib(0)=1,Fib(1)=1
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}

	return Fib(n-1) + Fib(n-2)
}

/*Fib2 using iteration*/
func Fib2(value int) int {

	x := make([]int, value)
	x[0], x[1] = 1, 1
	for n := 2; n < value; n++ {
		x[n] = x[n-1] + x[n-2]
	}
	return x[value-1]
}
