package main

import (
	"fmt"
)

func main() {
	// LABEL1:
	// 	for {
	// 		for i := 0; i < 10; i++ {
	// 			fmt.Println(i)
	// 			if i > 3 {
	// 				fmt.Println("jump LABEL1")
	// 				break LABEL1 // default break current for
	// 			}
	// 		}
	// 	}

	// 	for i := 0; i < 10; i++ {
	// 		if i > 5 {
	// 			goto LABEL3
	// 		}

	// 	}
	// LABEL3:
	// 	fmt.Println("LABEL3")

LABEL4:
	for i := 0; i < 10; i++ {
		fmt.Println(i)
		for {
			fmt.Println("continue LABEL4")
			continue LABEL4
		}

	}
	fmt.Println("OK")

}
