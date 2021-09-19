package main

import "fmt"

func bar() {
	xxx := func() map[int]func() int {
		var xs = map[int]func() int{
			1: func() int { return 10 },
			2: func() int { return 20 },
			3: func() int { return 30 }, //← Mandatory ,
		}
		return xs
	}
	fmt.Println(xxx())
}
