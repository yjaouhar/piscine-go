package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	sl := []rune(os.Args[0])
	for i := 2; i < len(sl); i++ {
		z01.PrintRune(sl[i])
	}
	z01.PrintRune('\n')
}
