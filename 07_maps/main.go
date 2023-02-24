package main

import (
	"fmt"
)

func main() {
	// maps
	// m := make(map[int]string)
	// m[1] = "shubam"
	// m[2] = "sharma"

	// declaring and adding key-value
	m := map[int]string{1: "shubam", 2: "Sharma", 3: "welcome"}

	// adding after declaring in map m
	m[4] = "home"

	//overriding the value on same key
	m[4] = "school"

	fmt.Println(m)
	fmt.Println(len(m))

	// Delete from map
	delete(m, 1)
	fmt.Println(m)
}
