package main

import "fmt"

func main() {
	switch {
	case false:
		fmt.Println("this should not print")
	case (2 == 4):
		fmt.Println("this should not print")
	case (3 == 3):
		fmt.Println("prints")
		fallthrough // allows the program to read the next case
	case (4 == 4):
		fmt.Println("also true, does it print?")
	}
}
