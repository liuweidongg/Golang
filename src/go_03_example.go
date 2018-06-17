package main

import (
	"fmt"
	"strconv"
)

func main() {
	var a int = 65
	b := string(a) // it's ASCII
	fmt.Println(b)

	b1 := strconv.Itoa(a) // string to int
	fmt.Println(b1)

	a, _ = strconv.Atoi(b1)
	fmt.Println(a)
}
