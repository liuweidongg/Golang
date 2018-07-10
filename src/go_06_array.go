package main

import (
	"fmt"
)

func main() {
	// define array
	// variable_name  [array_length] variable_type
	// var a [2]int
	// var b [1]int
	// diffrent array can't  assignment,array length belong type with Golang

	c := [2]int{1, 1}
	fmt.Println(c)
	d := [2]int{1}
	fmt.Println(d)
	d1 := [2]int{1: 1} // use index ensure elemnt position
	fmt.Println(d1)

	e := [...]int{1, 2, 3, 4, 5, 6} //Indefinite length array
	e1 := [...]int{0: 1, 1: 2, 2: 200}
	e2 := [...]int{10: 1}
	fmt.Println(e, e1, e2)

	var p *[11]int = &e2
	fmt.Println(p)

	x, y := 1, 2
	f := [...]*int{&x, &y} // return id in RAM, & :get id , * : pointer
	fmt.Println(f)

	// array == !=    just can use this
	g, h := [2]int{1, 2}, [2]int{1, 2} // note array_length
	fmt.Println(g == h)

	// new
	i := new([10]int)
	i[1] = 100
	fmt.Println(i) // return pointer array ,like var p *[11]int = &e2

	//  array operating
	j := [10]int{}
	j[1] = 2
	fmt.Println(j)

	// D-array
	k := [2][3]int{
		{1, 1, 1},
		{2, 2, 2},
	}
	fmt.Println(k)

	l := [...][3]int{
		{1, 1, 1},
		{2, 2, 2},
	}
	fmt.Println(l)

	//  maopao sorted

	m := [...]int{5, 2, 6, 3, 9}
	fmt.Println(m)
	num := len(m)
	for i := 0; i < num; i++ {
		for j := i + 1; j < num; j++ {
			if m[i] < m[j] {
				temp := m[i]
				m[i] = m[j]
				m[j] = temp
			}
		}
	}
	fmt.Println(m)
}
