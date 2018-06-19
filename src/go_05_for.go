package main

import (
	"fmt"
)

func main() {
	a := 1
	// modus 1
	/*
		for {
			a++
			if a > 3 {
				fmt.Println(a)
				break
			}
		}*/
	// modus 2
	/*
		for a < 10 {
				a++
				fmt.Println(a)
			}
			fmt.Println("Over")

	*/
	// don't put function in "for" condition,
	// because, it will be longer time
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
	fmt.Println("Over", a)

}
