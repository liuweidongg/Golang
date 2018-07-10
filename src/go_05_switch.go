package main

import (
	"fmt"
)

func main() {
	a := 1
	/*
		switch a {
			case 0:
				fmt.Println("a=0")
			case 1:
				fmt.Println("a=1")
			default:
				fmt.Println("None")
			}

	*/

	/*
	   switch {
	   	case a == 1:
	   		fmt.Println(a)
	   	case a == 2:
	   		fmt.Println(a)
	   	default:
	   		fmt.Println("None")

	   	}
	*/

	// a is local variable
	switch a := 1.0; { //constant 0.1 truncated to integer

	case a >= 0.0:
		fmt.Println("a>=0")
		fallthrough // continue check next case
	case a >= 0.1:
		fmt.Println("a >= 0.1")

	default:
		fmt.Println("None")
	}
	fmt.Println(a)
}
