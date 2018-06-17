package main

import (
	"fmt"
)

// const a int = 1
// const b  = "A"

// const (
// 	c = a
// 	d = a + 1
// 	e =  a + 2
// )
// const  a1,b1,c1   = 1,"2","C"

// const (
// 	a = 1
// 	b // b = a, if we not give  value of  b
// 	c
// )

// var sss = "123" // const initializer len(sss) is not a constant

// const (
// 	a = len(sss)
// 	b // b = a, if we not give  value of  b
// 	c
// )

// use const and  innfunctionn !!
// const (
// 	a = "123"
// 	b = len(a)
// 	c
// )

// const (
// 	a, b = 1, "2"
// 	c, d
// )

//enumerate
const (
	a = "A"
	b
	c = iota // every  start const the iota is 0,so ,a:0,b:1,c:2,d:3
	d
)
const (
	e = iota
)

func main() {
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)
}
