package main

import (
	"fmt"
)

/*
*Programmer: techideator@gmail.com

*Purpose:
 show variable declarations
 show conditions
 show looping

*/

func main() {
	x := 1.5
	fmt.Println(square(&x))

	/*x := []int{
		48, 96, 86, 68,
		57, 82, 63, 70,
		37, 34, 83, 27,
		19, 97, 9, 17,
	}
	fmt.Println(exercises.Min(x))
	*/
	//fmt.Print(exercises.BSort([]int{5, 2, 6, 19, 3}))
	//print(exercises.Fib2(9))
	//bar()
	//foo2(1, 23, 3, 5)
	//	rangeLoop()
	//exercises.PrintA()
	//exercises.ReplaceRune()
	//exercises.ReverseString("foobar")
	//exercises.CharacterCountInString()
}

func hello() {

	fmt.Println("Helo World")

	var msg = "Hello World"
	fmt.Printf(" %s \n", msg)

	var msg2 string = "Hello universe"
	fmt.Println(msg2)

	msg3 := "Hello Globe"
	fmt.Println(msg3)

	//////////////////////////////////////////////////////////////////////////////
	//Other var declarations and data types
	///////////////////////////////////////////////////////////////////////////

	var (
		name   string
		age    int
		single bool
	)

	name = "Ebite Zion"
	age = 24
	single = true
	//or name,age,single:= "ebite zion",23,true
	fmt.Printf("Name is : %s\n Age is: %d\n Single %t", name, age, single)

	//constant
	const PI = 3.142
	fmt.Printf("PI = %f", PI)

}

func loop() {
	count := 10
	for i := 0; i < count; i++ {
		fmt.Println(i)
	}
}

func multiVariateLoop() {
	count := 50
	for i, j := 0, 0; i < count; i, j = i+4, j+1 {
		fmt.Printf("%d--------------%d\n", i, j)
	}
}

func nestedLoop() {

	//With loops within loops you can specify a label after break. Making the label identify
	//which loop to stop:
J:
	for j := 0; j < 5; j++ {
		for i := 0; i < 10; i++ {
			if i > 5 {
				break J //← Now it breaks the j-loop, not the i one
			}
			println(i)
		}
	}

}

func rangeLoop() {
	//The keyword range can be used for loops. It can loop over slices, arrays, strings, maps
	//and channels
	list := []string{"a", "?", "w"}
	for k, v := range list {
		fmt.Printf("%s is at position %d \n", v, k)
	}
}

func unhexSwitch(c byte) byte {
	//There is no automatic fall through, you can however use fallthrough to do just that.
	//Without fallthrough:
	switch {
	case '0' <= c && c <= '9':
		return c - '0'
	case 'a' <= c && c <= 'f':
		return c - 'a' + 10
	case 'A' <= c && c <= 'F':
		return c - 'A' + 10
	}
	return 0
}
func shouldEscapeSwitch(c byte) bool {
	switch c {
	case ' ', '?', '&', '=', '#', '+': // ← , as ”or”
		return true
	}
	return false
}
func square(x *float64) float64 {
	*x = *x * *x
	return *x
}
