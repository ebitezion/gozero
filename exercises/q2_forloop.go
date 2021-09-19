package exercises

/*
 @author: techideator@gmail.com
Todo: 1. Create a simple loop with the for construct. Make it loop 10 times and print out
the loop counter with the fmt package.
2. Rewrite the loop from 1. to use goto. The keyword for may not be used.
3. Rewrite the loop again so that it fills an array and then prints that array to the
screen
*/
func forLoopTenTimes() int {
	count := 0
	for i := 0; i < 10; i++ {
		count++

	}
	return count
}

func gotoLoopTenTimes() int {
	//set count to be 0
	count := 0
REPEAT:
	//increase count by 1
	count = count + 1
	if count < 10 {
		goto REPEAT
	}

	return count
}

func LoopFillsArrayAndPrint() [10]int {
	//declare a size 10 array
	var arr [10]int

	//initialize the array
	for i := 0; i < 10; i++ {
		arr[i] = i
	}
	return arr
}
