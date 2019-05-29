package main

import "fmt"

func main() {
	foo()
	bar("William")
	s1 := woo("Moneypenny")
	fmt.Println(s1)
	x, y := mouse("Ian", "Fleming")
	fmt.Println(x)
	fmt.Println(y)
}

func foo() {
	fmt.Println("Hello from foo")
}

// everything in Go is PASS BY VALUE
func bar(s string) {
	fmt.Println("Hello,", s)
}

// single return
func woo(st string) string {
	return fmt.Sprint("Hello from woo, ", st)
}

// multiple returns
func mouse(fn string, ln string) (string, bool) {
	a := fmt.Sprint(fn, " ", ln, "says hello")
	b := false
	return a, b
}
