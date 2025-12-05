package main

import "github.com/01-edu/z01"

var b string = "x = 42, y = 21"

func main() {
	for _, ch := range b {
		z01.PrintRune(ch)
	}
	z01.PrintRune('\n')
}
