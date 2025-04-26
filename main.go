package main

import (
	"fmt"
	"strings"
)

func runPart1() {
	fmt.Println(CaesarCipher("I love CS!", 5), "\nExpected: NQTAJHX")
	fmt.Println(CaesarCipher("Hello, World!", 3), "\nExpected: KHOORZRUOG")
	fmt.Println(CaesarCipher("ABC XYZ", 2), "\nExpected: CDEZAB")
	fmt.Println(CaesarCipher("Go Concurrency", 7), "\nExpected: NVJVTULYNLZAPJP")
	fmt.Println(CaesarCipher("SYSC3101!", 4), "\nExpected: WCMG")
}

func runPart2() {
	concurrent()
}

func runPart3() {
	accelerateConcurrency()
}

func main() {
	fmt.Println(strings.Repeat("-", 25))
	fmt.Println("Part 1")
	runPart1()
	fmt.Println(strings.Repeat("-", 25))
	fmt.Println("Part 2")
	runPart2()
	fmt.Println(strings.Repeat("-", 25))
	fmt.Println("Part 3")
	runPart3()
}
