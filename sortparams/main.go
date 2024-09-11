package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1:] // excluding the program name itself

	// Bubble sort algorithm to sort the arguments in ASCII order
	for i := 0; i < len(args)-1; i++ {
		for j := 0; j < len(args)-1-i; j++ {
			if args[j] > args[j+1] {
				args[j], args[j+1] = args[j+1], args[j]
			}
		}
	}

	// Print the sorted arguments
	for _, arg := range args {
		for _, ch := range arg {
			z01.PrintRune(ch)
		}
		z01.PrintRune('\n')
	}
}
