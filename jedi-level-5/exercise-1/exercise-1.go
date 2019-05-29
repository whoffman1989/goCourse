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

	fmt.Println(p1.firstName)
	fmt.Println(p1.lastName)
	fmt.Println("Favorite Ice Cream Flavors:")
	for i, v := range p1.iceCreamFlavors {
		fmt.Println(i, v)
	}

	fmt.Println(p2.firstName)
	fmt.Println(p2.lastName)
	fmt.Println("Favorite Ice Cream Flavors:")
	for i, v := range p2.iceCreamFlavors {
		fmt.Println(i, v)
	}
}
