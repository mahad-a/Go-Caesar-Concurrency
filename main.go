package main

import (
	"fmt"
	"strings"
)

func runQuestion1() {
	fmt.Println(CaesarCipher("I love CS!", 5), "\nExpected: NQTAJHX")
	fmt.Println(CaesarCipher("Hello, World!", 3), "\nExpected: KHOORZRUOG")
	fmt.Println(CaesarCipher("ABC XYZ", 2), "\nExpected: CDEZAB")
	fmt.Println(CaesarCipher("Go Concurrency", 7), "\nExpected: NVJVTULYNLZAPJP")
	fmt.Println(CaesarCipher("SYSC3101!", 4), "\nExpected: WCMG")
}

func runQuestion2() {
	concurrent()
}

func runQuestion3() {
	accelerateConcurrency()
}

func main() {
	fmt.Println(strings.Repeat("-", 25))
	fmt.Println("Question 1")
	runQuestion1()
	fmt.Println(strings.Repeat("-", 25))
	fmt.Println("Question 2")
	runQuestion2()
	fmt.Println(strings.Repeat("-", 25))
	fmt.Println("Question 3")
	runQuestion3()
}
