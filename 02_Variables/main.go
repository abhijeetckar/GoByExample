package main

import "fmt"

func main() {

    var a, name string = "initial" , "Abhijeet"
    fmt.Println(a)
    fmt.Println(name)

    var b, c int = 1, 2
    fmt.Println(b, c)

    var d = true
    fmt.Println(d)

    var e int
    fmt.Println(e)

    f, g := "apple", "orange"
    fmt.Println(f)
    fmt.Println(g)

	var h, i = 6, "Hello"
  	j, k := 7, "World!"

	fmt.Println(h)
	fmt.Println(i)
	fmt.Println(j)
	fmt.Println(k)

	var (
		l int
		m int = 1
		n string = "hello"
	  )

	fmt.Println(l)
	fmt.Println(m)
	fmt.Println(n)


	//redeclaration of vars
	aa, bb := 1, 2
    bb, cc := 3, 4
    fmt.Println(aa, bb, cc)
}