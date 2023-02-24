package main

import "fmt"

// return type
func add(x int, y int) int {
	return (x + y)
}

// without return type 
func greeting(name string) {
	fmt.Println("hello", name)
}
func main() {
	fmt.Println(add(3, 5))
	greeting("shubam")
}
