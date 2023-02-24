package main

import "fmt"

func main() {
	x := 5
	y := 10

	if x < y {
		fmt.Printf("%d is less than %d\n", x, y)
	} else {
		fmt.Printf("%d is greater than %d\n", y, x)
	}

	// switch statement
	number := 3

	// Switch
	switch number {
	case 1:
		fmt.Println("you choose 1")
	case 2:
		fmt.Println("you choose 2")
	default:
		fmt.Println("you choose other than 1 and 2")
	}

	// sum using loop
	sum := 0
	for i := 0; i < 10; i++ {
		sum = sum + i
	}
	fmt.Println("Sum is ", sum)

}
