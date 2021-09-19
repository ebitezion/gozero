package exercises

/*Q7. (0) Integer ordering
1. Write a function that returns its (two) parameters in the right, numerical (ascending) order:
f(7,2) → 2,7
f(2,7) → 2,7
*/
func f(x, y int) (int, int) {
	if x > y {
		x, y = y, x
		return x, y
	}

	return x, y
}
