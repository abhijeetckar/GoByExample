package main

import "fmt"

func main()  {
	name := "Abhijeet"

	for i := 0; i < len(name); i++ {
		// fmt.Println(string(name[i]))
		// fmt.Println(len(name))
	}

	for _, v := range name {
		// fmt.Println(_)
		fmt.Println(string(v))
	}

	for {
        fmt.Println("loop")
        break
    }

	for n := 0; n <= 5; n++ {
        if n%2 == 0 {
            continue
        }
        fmt.Println(n)
    }
}