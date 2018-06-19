package main

import (
	"fmt"
)

func main() {
	a := 2
	if a > 1 {
		fmt.Println("OK")
	}

	if b := 1; b == 1 { // can use ,if b in "if"
		fmt.Println("OK B")
	}
	// fmt.Println(b)  can not get b

	if a := 10; a == 10 {
		fmt.Println(a)
	} // use a in if, not a :=2

	if a, b := 10, 11; a+b == 21 {
		fmt.Println("OKOK")
	}
}
