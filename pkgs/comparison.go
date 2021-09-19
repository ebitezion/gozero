package pkgs

/*
*Programmer: techideator@gmail.com

*Purpose:
 shows comparison functions
*/

func compare(a, b []byte) int {
	/*
			 Compare returns an integer comparing the two byte arrays lexicographically. The
			 result will be 0 if a == b, -1 if a < b, and +1 if a > b ;

			 Strings are equal except for possible tail;
		      Strings are equal.
	*/

	for i := 0; i < len(a) && i < len(b); i++ {
		switch {
		case a[i] > b[i]:
			return 1
		case a[i] < b[i]:
			return -1
		}
	}

	switch {
	case len(a) < len(b):
		return -1
	case len(a) > len(b):
		return 1
	}
	return 0

}
