package main

import (
	"fmt"
	"sort"
)

func main() {
	slc := []int{3, 4, 2, 1}

	// sort.Slice(slc, func(i, j int) bool {
	// 	return slc[i] < slc[j]
	// })

	fmt.Println(slc)

	slc2 := []int{4, 2, 3, 1, 5, 7, 11, 9, 8}

	sort.Slice(slc2[2:5], func(i, j int) bool {
		return slc2[2+i] < slc2[2+j]
	})

	fmt.Println(slc2)
}
