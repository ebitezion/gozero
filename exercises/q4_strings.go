package exercises

import "fmt"

/*
Q4. (1) Strings
1. Create a Go program that prints the following (up to 100 characters):
A
AA
AAA
AAAA
AAAAA
AAAAAA
AAAAAAA
...
2. Create a program that counts the number of characters in this string:
asSASA ddd dsjkdsjs dk
In addition, make it output the number of bytes in that string. Hint: Check out the
utf8 package.
3. Extend/change the program from the previous question to replace the three runes
at position 4 with ’abc’.
4. Write a Go program that reverses a string, so “foobar” is printed as “raboof”. Hint:
You will need to know about conversion; skip ahead to section “Conversions” on
page 54.
*/

//Executed in the main method of main.go
func PrintA() {
	/* if arrays were used the ff is ok
	var arr []string
	for i := 0; i < 100; i++ {
		arr = append(arr, "A")
		fmt.Println(arr)
	}*/
	//requires us to use string
	var str string = "A" //strings are immutable
	//convert strings to runes
	characters := []rune(str)

	for i := 0; i < 100; i++ {
		fmt.Println(string(characters))
		characters = append(characters, 'A')

	}

}

func CharacterCountInString() {

	str := "asSASA ddd dsjkdsjs dk"

	var a, d, j, k, s, A, S int

	for _, v := range str {

		if v == 'a' {
			a++
		}

		if v == 'd' {
			d++
		}

		if v == 'j' {
			j++
		}

		if v == 'k' {
			k++
		}

		if v == 's' {
			s++
		}

		if v == 'S' {
			S++
		}

		if v == 'A' {
			A++
		}

	}
	fmt.Printf("a %d\n d %d\n s %d\n k %d\n S %d\n A %d\n j %d\n", a, d, s, k, S, A, j)

}
func ReplaceRune() {
	/*possible solution
	      s := "ðåぽ ai no Æl"
	r := []rune(s)
	copy(r[4:4+3], []rune("abc"))
	fmt.Printf("Before: %s\n", s);
	fmt.Printf("After : %s\n", string(r))
	*/
	str := "asSASA ddd dsjkdsjs dk"
	obj := []rune(str)

	obj[4] = 'a'
	obj[5] = 'b'
	obj[6] = 'c'
	fmt.Println(string(obj))

}

func ReverseString(str string) {
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
	fmt.Println(string(newChar))
}
