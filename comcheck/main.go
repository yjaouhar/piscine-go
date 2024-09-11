package main

import (
	"fmt"
	"os"
)

func main() {
	sl := []string(os.Args)
	for _, v := range sl {
		if v == "01" || v == "galaxy" || v == "galaxy 01" {
			fmt.Println("Alert!!!")
			break
		}
	}
}
