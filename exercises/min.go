/*
4. Write a program that finds the smallest number
in this list:
x := []int{
 48,96,86,68,
 57,82,63,70,
 37,34,83,27,
 19,97, 9,17,
}
	//what is the lowest int available? in 32 bits os
	//max := int(^uint(0) >> 1) or min:=-min-1 or minint32=-1<<31-1 or but better to use math pkg
	//min32 := -1<<31 - 1
*/package exercises

import (
	"errors"
	"reflect"
	"strings"
)

func Min(x []int) (int, error) {

	err := sanitize(x)
	if err != nil {
		return -1, err
	}

	var smallest int = x[0]
	//check each index for number min number
	//ls := BSort(x)
	//fmt.Println(ls)

	for k, _ := range x {

		if x[k] < smallest {
			//this is the new min
			smallest = int(x[k])
		} else {
			smallest = smallest
		}

	}
	//check for errors-- usually comes from user

	return smallest, err
}
func sanitize(x []int) error {
	//check input
	emptyErr := errors.New("Input: Empty")
	invalidTErr := errors.New(strings.ToLower("Input: Invalid Expected a []int slice"))
	incErr := errors.New("Input: Incomplete slice,please give two or more")

	if len(x) == 1 {
		return incErr
	}
	if x == nil {
		return emptyErr
	}
	var i []int
	if reflect.TypeOf(x) != reflect.TypeOf(i) {
		return invalidTErr
	}

	return nil
}
