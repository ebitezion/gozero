package pkgs

/*
*Programmer: techideator@gmail.com

*Purpose: Show how to work with strings

 */

func mutatingStr(str string) string {
	//change the first character of str
	character := []rune(str)

	character[0] = 'D'

	newChar := string(character)
	return newChar
}

func ReverseString(str string) string {
	/*
	   . Reversing a string can be done as follows. We start from the left (i) and the right
	   (j) and swap the characters as we see them:
	   Listing 1.11. Reverse a string
	   import "fmt"
	   func main() {
	   s := "foobar"
	   a := []rune(s) //←  a conversion
	   for i, j := 0, len(a)-1; i < j; i, j = i+1, j-1 {
	   a[i], a[j] = a[j], a[i] ← Parallel assignment
	   }
	   fmt.Printf("%s\n", string(a)) ← Convert it back
	   }
	*/

	//conv to rune
	characters := []rune(str)
	//range through the length, from max to 0, each time decrementing by -1, like n-1.
	var newChar []rune
	var n int = len(characters)
	for {

		if n == 0 {
			break
		}

		newChar = append(newChar, characters[n-1])
		n--
	}
	return string(newChar)
}
