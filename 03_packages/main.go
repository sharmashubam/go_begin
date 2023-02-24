package main

import (
	"fmt"
	"math"

	"github.com/sharmashubam/go_begin/01_hello/03_packages/strutil"
)

func main() {
	fmt.Println(math.Floor(4.3344))
	fmt.Println(math.Ceil(2.33))
	fmt.Println(strutil.Reverse("hello"))
}
