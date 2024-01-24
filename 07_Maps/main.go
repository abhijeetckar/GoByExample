package main

import "fmt"

func main() {
	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("map:", n)

	m := make(map[string]string)

	m[`k1`] = `a`
	m[`k2`] = `b`

	fmt.Println(m)

	val, ok := m[`k1`]
	fmt.Println(val) //val
	fmt.Println(ok)  // true

	k1, v1 := m[`k3`]
	fmt.Println(k1) //prints nothing
	fmt.Println(v1) //false

	for k, v := range m {
		fmt.Println(k)
		fmt.Println(v)
	}

	maps := make([]map[string]string, 3)

	map1 := make(map[string]string)
	map3 := make(map[string]string)
	map2 := make(map[string]string)
	map1[`a`] = `1`
	map2[`b`] = `2`
	map3[`c`] = `3`

	maps[0] = map1
	maps[1] = map2
	maps[2] = map3

	for _, v := range maps {
		fmt.Println(v)
	}
}
