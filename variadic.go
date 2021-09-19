package main

import (
	"fmt"
)

func foo(args ...int) {

	fmt.Println(args[:1])
}

func foo2(args ...int) {
	foo(args...)
}
