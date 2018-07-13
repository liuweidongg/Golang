package main

import (
	"fmt"
	"sort"
)

func main() {
	// the map is disorder!!!
	var m map[int]string // int == key ,string == value   map: key-value
	// m = map[int]string{}
	// fmt.Println(m)

	// make, must initial !!!
	m = make(map[int]string)
	fmt.Println(m)

	m1 := make(map[int]string)
	fmt.Println(m1)

	// map like dict in Python
	m1[1] = "OK"
	m1[2] = "OJBK"
	// get value
	a := m1[1]
	// not'd crash,return empty, if the key not in this  map,
	// and get empty if u not set key-value in this map
	b := m1[30]

	// delete key-value
	delete(m1, 1)
	fmt.Println(m1)
	fmt.Println(a)
	fmt.Println(b)

	// Multistage map, must initial map ,

	var m2 map[int]map[int]string
	m2 = make(map[int]map[int]string) // initial first map
	m2[1] = make(map[int]string)
	m2[1][1] = "PK"
	c := m2[1][1]
	fmt.Println(c)
	fmt.Println(m2)

	d, ok := m2[1][1] //if ok is True ,it's  mean is initial for  this map ,else ,not intial.
	fmt.Println(d, ok)

	// ****
	e, ok2 := m2[2][2]
	fmt.Println(e, ok2)
	if !ok2 {
		m2[2] = make(map[int]string)
	}
	m2[2][2] = "Good"
	e, ok2 = m2[2][2]
	fmt.Println(e, ok2)

	sm := make([]map[int]string, 5)
	// for _, v := range sm {
	// 	v = make(map[int]string, 1)
	// 	v[1] = "OK" // it's v is copy element ,so ,it cand'nt modifily sm
	// 	fmt.Println(v)
	// }
	// fmt.Println(sm)
	for i := range sm {
		sm[i] = make(map[int]string, 1)
		sm[i][1] = "OK" // it's v is copy element ,so ,it cand'nt modifily sm
		fmt.Println(sm[i])
	}
	fmt.Println(sm)

	//envent the map is disorder but we can use key

	tm := map[int]string{1: "a", 2: "b", 3: "c", 4: "d"}
	s := make([]int, len(tm)) // it's slice
	i := 0
	for k, _ := range tm {
		s[i] = k // the k is disorder ,so this slice as so as k
		i++
	}
	sort.Ints(s)
	fmt.Println(s)

}
