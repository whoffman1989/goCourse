package main

import "fmt"

func main() {
	fmt.Println("Hello everyone.")

	foo()

	fmt.Println("Hello again.")

	for i := 0; i < 100; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		}
	}
}

func foo() {
	fmt.Println("I am in foo.")
}
