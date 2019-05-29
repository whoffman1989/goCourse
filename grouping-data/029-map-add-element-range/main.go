package main

import "fmt"

func main() {
	m := map[string]int{
		"James":           32,
		"Miss Moneypenny": 27,
	}
	fmt.Println(m)

	m["William"] = 29

	for k, v := range m {
		fmt.Println(k, v)
	}
}
