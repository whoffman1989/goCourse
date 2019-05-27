package main

import "fmt"

func main() {
	m := map[string][]string{
		"bond_james":      []string{"shaken, not stirred", "martinis", "women"},
		"moneypenny_miss": []string{"james bond", "literature", "computer science"},
		"no_dr":           []string{"being evil", "ice cream", "sunsets"},
	}
	// fmt.Println(m)

	m["flemming_ian"] = []string{"steaks", "cigars", "espionage"}

	delete(m, "no_dr")

	for k, v := range m {
		fmt.Println("This is the record for: ", k)
		for i, v2 := range v {
			fmt.Println("\t", i, v2)
		}
	}
}
