package main

import (
	"fmt"
)

func main() {
	var s1 []int
	fmt.Println(s1)
	a := [10]int{1, 2, 3, 4, 5, 6, 8, 9, 10}
	fmt.Println(a)
	s2 := a[5:10]
	fmt.Println(s2)
	s3 := a[5:len(a)] // a[5:]     a[:5]
	fmt.Println(s3)

	s4 := make([]int, 3, 10) // tyep int, 3 elements in s4 ,10 length RAM,but can explore duble.
	fmt.Println(len(s4), cap(s4))

	s5 := make([]int, 3) // cap is 3
	fmt.Println(len(s5), cap(s5))
	b := []byte{'a', 'b', 'c', 'd', 'e'}
	fmt.Println(b)
	sb := b[1:]
	fmt.Println(len(sb), cap(sb)) // 4 elements in sb, and 4 length  RAM
	sb2 := sb[0:2]                // reslice
	fmt.Println(string(sb))
	fmt.Println(string(sb2))

	// append
	s6 := make([]int, 3, 6)
	fmt.Printf("%p\n", s6)
	s6 = append(s6, 1, 2, 4)
	fmt.Printf("%v  %p\n", s6, s6)
	s6 = append(s6, 1, 2, 4) // if append elements length > inital slice ,than,will create new slice in RAM,and copy inital slice to new slice
	fmt.Printf("%v  %p\n", s6, s6)

	// importent
	c := []int{1, 2, 3, 4, 5}
	s7 := c[2:5]
	s8 := c[1:3]
	fmt.Println(s7, s8)
	// the s7 and s8 pointer the "3"
	// s8 = append(s8, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1)
	// if elements length > inital slice ,than s8's RAM will be change ,so now ,we change s7,the s8 is not change
	s7[0] = 9
	fmt.Println(s7, s8)

	// copy
	s9 := []int{1, 2, 3, 4, 5, 6}
	s10 := []int{7, 8, 9}
	// copy(s9, s10) //   [7 8 9 4 5 6],
	// copy(s9, s10[0:2]) // partial copy
	copy(s9[0:1], s10[0:1])
	fmt.Println(s9)

}
