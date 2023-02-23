package main

import "fmt"

func main() {
	// MAIN TYPES
	// string
	// bool
	// int
	// int  int8  int16  int32  int64
	// uint uint8 uint16 uint32 uint64 uintptr
	// byte - alias for uint8
	// rune - alias for int32
	// float32 float64
	// complex64 complex128

	// short hand method of declaring the variables
	message := "hello"
	name, email := "shubam", 12

	// using var
	var age int32 = 17
	fmt.Println(age, message, email, name)

	// using float and format verbs
	var scien = -1.234456e+78
	fmt.Printf("%e\n", scien)

	// getting the type
	fmt.Printf("%T\n", name)
	fmt.Printf("Type is: %T\n", scien)

	// using const
	const isCool = true
	// isCool = false
	fmt.Println(isCool)

}
