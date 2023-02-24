package main

import (
	"fmt"
)

func main() {
	ids := []int{32, 43, 23, 56, 4, 65, 34}

	for i, id := range ids {
		fmt.Println(i, "matches", id)
	}

	// if we dont want to have index

	for _, id := range ids {
		fmt.Println(id)
	}

	//  getting sum
	sum := 0
	for _, id := range ids {
		sum = sum + id
	}
	fmt.Println("Sum", sum)

	// range with map
	emails := map[string]string{"Bob": "bob@gmail.com", "Sharon": "sharon@gmail.com"}

	for k, v := range emails {
		fmt.Println(k, "has value", v)
	}

	for k := range emails {
		fmt.Println("Name: " + k)
	}

	for _, val := range emails {
		fmt.Println("emails: " + val)
	}

}
