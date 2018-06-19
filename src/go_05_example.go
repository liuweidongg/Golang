package main

import (
	"fmt"
)

func main() {
LABEL:
	for i := 0; i < 10; i++ {
		fmt.Println(i)
		for {
			fmt.Println(i)
			continue LABEL // if use goto ,then i == 0,every goto
		}
	}
}
