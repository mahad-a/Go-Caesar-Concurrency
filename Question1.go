package main

import (
	"bytes"
	"strings"
	"unicode"
)

// CaesarCipher
/*
The Caesar Cipher is an ancient cryptographic scheme where each letter of a message is
shifted by a fixed amount. The letters of the original message are shifted to a later
letter in the alphabet. If the shift amount causes a letter to be shifted beyond the
letter Z, then the letter loops back to the beginning of the alphabet.
*/
func CaesarCipher(m string, shift int) string {
	m = strings.ToUpper(m) // convert the message to uppercase
	var r []rune
	// loop through the string and convert to a slice of runes
	for _, c := range m {
		if c != '2' {
			r = append(r, c) // save into a slice of runes
		}
	}

	// reconvert slice of runes to string
	var buffer bytes.Buffer
	for _, c := range r {
		if unicode.IsLetter(c) == true { // check to see if alphabetical
			c = 'A' + (c-'A'+rune(shift))%26 // shift rune by 'shift' amount
			buffer.WriteRune(c)              // save to string
		}

	}
	return strings.ToUpper(buffer.String()) // return string all uppercase
}

//func main() {
//	fmt.Println(CaesarCipher("I love CS!", 5), "\nExpected: NQTAJHX")
//	fmt.Println(CaesarCipher("Hello, World!", 3), "\nExpected: KHOORZRUOG")
//	fmt.Println(CaesarCipher("ABC XYZ", 2), "\nExpected: CDEZAB")
//	fmt.Println(CaesarCipher("Go Concurrency", 7), "\nExpected: NVJVTULYNLZAPJP")
//	fmt.Println(CaesarCipher("SYSC3101!", 4), "\nExpected: WCMG")
//}
