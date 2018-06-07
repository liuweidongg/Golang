package main

import (
	"fmt"
	"math"
)

//Alias
type (
	byte int8
	rune int32
	文本   string // like Python unicode variable name
)

func main() {
	var a int
	fmt.Println(a)
	var b float32
	fmt.Println(b)
	var c bool
	fmt.Println(c)
	var d string
	fmt.Println(d)
	var e [10]int
	fmt.Println(e)            //  array type is int,10 is size
	fmt.Println(math.MaxInt8) // max int8 value

	var b1 文本 // prohibit same name !!!
	b1 = "chinese text name"
	fmt.Println(b1)

	var b2 int = 100
	fmt.Println(b2)

	var b3 = "Joker" // like JS variable
	fmt.Println(b3)

	b4 := 100.01 // simple var
	fmt.Println(b4)

	var a1, b5, c1, d1 = 1, 2, 3, "joker" //local variable ,var () global variable!!
	fmt.Println(a1, b5, c1, d1)

	a11, _, c11, d11 := 1, 2, 3, 4
	fmt.Println(a11, c11, d11) // '_' means ignore, and can't use _ as value

	// type change
	var a22 float32 = 1.1
	bb := int(a22)

	// cc := true
	// dd := int(cc)  can't change type,bool just have true false
	// var str = "123"
	// ss := int(str)

}
