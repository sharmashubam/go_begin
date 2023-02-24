package main

import "fmt"

func main() {

	// arrays have fixed sizes
	var a [2]string
	a[0] = "Hello"
	a[1] = "World"

	fmt.Println(a[0], a[1])
	fmt.Println(a)

	// shorthand way of declaring arrrays
	fruits := [3]string{"apple", "mango", "grapes"}
	fmt.Println(fruits)

	// slice, on the other hand, is a dynamically-sized, flexible view into the elements of an array
	mySlice := []string{"shubam", "sharma", "hello", "welcome"}
	fmt.Println(mySlice)

	// functions on arrays
	// cap() function - returns the capacity of the slice (the number of elements the slice can grow or shrink to)

	fmt.Println(len(mySlice))
	fmt.Println(cap(mySlice))
}
