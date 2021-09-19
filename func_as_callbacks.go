package main

import "fmt"

func printsmth(no int) {
	fmt.Println(no)
}

func callback(y int, f func(int)) {

	f(y)
}
