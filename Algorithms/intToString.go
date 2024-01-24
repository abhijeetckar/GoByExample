package main

import (
	"fmt"
	"strconv"
)

func main() {
	slc := []int{3, 1, 2}

	str := ``
	for _, v := range slc {
		str += strconv.Itoa(v)
	}
	fmt.Println(str)
}
