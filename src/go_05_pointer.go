package main

import (
	"fmt"
)

func main() {
	/*
		"&" get variable id in RAM
		"*" use id get variable

	*/
	a := 1
	var p *int = &a // id in RAM
	fmt.Println(p)
	fmt.Println(*p)

	b := 1
	b++ // b--
	var p1 *int = &b
	fmt.Println(p1)
	fmt.Println(*p1)
}
