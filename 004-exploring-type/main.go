package main

import "fmt"

var y = 42

// DECLARE the VARIABLE with the INDENTIFIER "z"
// is of TYPE string
// and I ASSIGN the VALUE "shaken, not stirred"

var z = "shaken, not stirred"
var a = `James said, "Shaken, not stirred"`

// GO is a STATIC programming language
// a VARIABLE is DECLARED to hold a VALUE of a certain TYPE
// not a DYNAMIC programming language

func main() {
	fmt.Println(y)
	fmt.Printf("%T\n", y)
	fmt.Println(z)
	fmt.Printf("%T\n", z)
	fmt.Println(a)
	fmt.Printf("%T\n", a)
}
