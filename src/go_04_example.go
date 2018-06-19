package main

import (
	"fmt"
)

const (
	B  float64 = 1 << (iota * 10) // iota == 0
	KB                            // iota == 1
	MB
	GB
	TB
	PB
	EB
	ZB
	YB
)

func main() {
	// 1 -> b
	fmt.Println(1 << 10)       // b -> kb
	fmt.Println(1 << 10 << 10) // kb -> mb

	fmt.Println("---------------")
	fmt.Println(B)
	fmt.Println(KB)
	fmt.Println(MB)
	fmt.Println(GB)
	fmt.Println(TB)
	fmt.Println(PB)
	fmt.Println(EB)
	fmt.Println(ZB)
	fmt.Println(YB)
}
