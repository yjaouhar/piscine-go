package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("File name missing")
		return
	} else if len(os.Args) > 2 {
		fmt.Println("Too many arguments")
	} else {
		filename := os.Args[1]

		content, err := os.ReadFile(filename)
		if err != nil {
			return
		}

		fmt.Print(string(content))
	}
}
