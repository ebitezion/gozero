/*Package exercises Q14. (1) Bubble sort
1. Write a function that performs a bubble sort on a slice of ints. From [24]:
It works by repeatedly stepping through the list to be sorted, comparing
each pair of adjacent items and swapping them if they are in the wrong
order. The pass through the list is repeated until no swaps are needed,
which indicates that the list is sorted. The algorithm gets its name from
the way smaller elements “bubble” to the top of the list.
[24] also gives an example in pseudo code:
procedure bubbleSort( A : list of sortable items )
do
swapped = false
for each i in 1 to length(A) - 1 inclusive do:
if A[i-1] > A[i] then
swap( A[i-1], A[i] )
swapped = true
end if
end for
while swapped
end procedure*/package exercises

// BSort ...a func whose purpose in life is a sort and integer slice
var swapped = true

func BSort(A []int) []int {

	for swapped {
		swapped = false
		for i := 1; i < len(A); i++ {
			if A[i-1] > A[i] {
				A[i-1], A[i] = A[i], A[i-1]

				swapped = true
			}

		}

	}
	return A
}
