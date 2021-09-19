package exercises

import "strconv"

/*
Q9. (1) Stack
1. Create a simple stack which can hold a fixed number of ints. It does not have to
grow beyond this limit. Define push – put something on the stack – and pop –
retrieve something from the stack – functions. The stack should be a LIFO (last
in, first out) stack
*/

type stack struct { // ← stack is not exported
	i    int
	data [10]int
}

/*
func Pop(arr []int) []int {
	//retrieve last index and delete

	var newArr []int = arr[:len(arr)-1]
	return newArr
}
func Push(arr []int, item int) []int {
	arr = append(arr, item)
	return arr
}
*/
func (s *stack) push(k int) {
	s.data[s.i] = k
	s.i++
}
func (s *stack) pop() int {
	s.i--
	return s.data[s.i]
}
func (s stack) String() string {
	var str string
	for i := 0; i <= s.i; i++ {
		str = str + "[" + strconv.Itoa(i) + ":" + strconv.Itoa(s.data[i]) + "]"
	}
	return str
}
