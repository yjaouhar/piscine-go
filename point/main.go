package main

import "github.com/01-edu/z01"

func conv(n int) {
	c := '0'
	if n < 10 {
		for i := 0; i < n; i++ {
			c++
		}
		z01.PrintRune(c)
	} else {
		conv(n / 10)
		conv(n % 10)
	}
}

type point struct {
	x int
	y int
}

func setPoint(ptr *point) {
	ptr.x = 42
	ptr.y = 21
}

func main() {
	points := &point{}

	setPoint(points)

	print_x := "x = "
	print_y := ", y = "
	for _, vx := range print_x {
		z01.PrintRune(vx)
	}
	conv(points.x)
	for _, vy := range print_y {
		z01.PrintRune(vy)
	}
	conv(points.y)
	z01.PrintRune('\n')
}
