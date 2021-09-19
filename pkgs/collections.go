package pkgs

/*
*Programmer: techideator@gmail.com

*Purpose:
 shows collections
*/

func getIndex(i int) int {
	//using array
	var arr [10]int
	arr[0] = 42
	arr[1] = 13
	return arr[i]
}

func getSlice(n, m int) (slice []string, err error) {
	// using slices
	ASlice := []string{"Hello", "World", "Global", "Universe"}
	if err != nil {
		return nil, err
	}
	return ASlice[n:m], nil
}
func appendToSlice() []int {
	s0 := []int{0, 0}
	s1 := append(s0, 2)
	s2 := append(s1, 3, 5, 7)
	s3 := append(s2, s0...)
	//.0 append a single element, s1 == []int{0, 0, 2};
	//.1 append multiple elements, s2 == []int{0, 0, 2, 3, 5, 7};
	//.2 append a slice, s3 == []int{0, 0, 2, 3, 5, 7, 0, 0}. Note the three dots!
	//And
	//The function copy copies slice elements from a source src to a destination dst
	//and returns the number of elements copied. Source and destination may overlap.
	//The number of elements copied is the minimum of len(src) and len(dst).
	return s3
}
func copyToSlice() []int {
	var a = [...]int{0, 1, 2, 3, 4, 5, 6, 7}
	var s = make([]int, 6)
	copy(s, a[0:]) //← n1 == 6,s == []int{0, 1, 2, 3, 4, 5}
	copy(s, s[2:]) //← n2 == 4, s == []int{2, 3, 4, 5, 4, 5}

	return s
}
func aMap() map[string]int {
	//demo of a map
	monthdays := map[string]int{
		"Jan": 31, "Feb": 28, "Mar": 31,
		"Apr": 30, "May": 31, "Jun": 30,
		"Jul": 31, "Aug": 31, "Sep": 30,
		"Oct": 31, "Nov": 30, "Dec": 31, // ← Comma required
	}
	return monthdays
}
