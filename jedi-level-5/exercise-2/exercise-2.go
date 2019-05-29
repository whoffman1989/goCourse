package main

import "fmt"

type person struct {
	firstName       string
	lastName        string
	iceCreamFlavors []string
}

func main() {
	p1 := person{
		firstName: "William",
		lastName:  "Hoffman",
		iceCreamFlavors: []string{
			"Chocolate",
			"Vanilla",
			"cappuccino",
		},
	}
	p2 := person{
		firstName: "Lexi",
		lastName:  "Hoffman",
		iceCreamFlavors: []string{
			"Strawberry",
			"Pistachio",
			"martini",
		},
	}

	m := map[string]person{
		p1.lastName: p1,
		p2.lastName: p2,
	}

	for _, v := range m {
		fmt.Println(v.firstName)
		fmt.Println(v.lastName)
		for i, val := range v.iceCreamFlavors {
			fmt.Println(i, val)
		}
		fmt.Println("-----")
	}
}
