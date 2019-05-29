package main

import "fmt"

func main() {
	// by declaring x:=42; that narrows the scope to the if statement
	if x := 42; x == 42 {
		fmt.Println("001")
	}
}
