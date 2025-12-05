package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1:]
	for i := len(args) - 1; i >= 0; i-- {
		word := args[i]
		for _, ch := range word {
			z01.PrintRune(ch)
		}
		z01.PrintRune('\n')

	}
}
