package main

import "fmt"

func main() {
	a := 12
	b := &a

	fmt.Println(a, b)
	fmt.Printf("%T\n", b)

	//  Use * to read val from address
	fmt.Println(*b)
	fmt.Println(*&a)

	// changing val with pointer --> changed the value at the address
	*b = 10
	fmt.Println(a)
}