package main

import "fmt"

func main() {
	//var name []type
	var slice_type []string
	fmt.Println(slice_type)

	//Array
	var arr = [5]int{1, 2, 3, 4, 5}
	fmt.Println(len(arr))

	//Array
	arr2 := [...]string{"Abhijeet", "Tanmay"}
	fmt.Println(cap(arr2))
	fmt.Println(arr2[0])

	//Slice
	slc := make([]int, 2)
	fmt.Println(slc)
	fmt.Println(len(slc))
	fmt.Println(cap(slc))

	slc2 := arr[1:4]
	fmt.Println(slc2)
	fmt.Println(cap(slc))
	slc3 := append(slc2, 6, 7, 8)
	fmt.Println(slc3)
	fmt.Println(cap(slc3))

	//how does slice capacity work?
	//https://go.dev/blog/slices-intro

	//find an element in an array or slice

	var testArr = []string{`abhi`, `tanu`, `priti`, `vibha`, `n`, `s`}
	result := find(testArr, `aabhi`)
	fmt.Println(result)
	fmt.Println(len(testArr))
	fmt.Println(cap(testArr))

	resultFnD := findAndDelete2(testArr, `abhi`)
	fmt.Println(resultFnD)
	fmt.Println(len(resultFnD))
	fmt.Println(cap(resultFnD))

}

func find(arr []string, item string) bool {
	for _, v := range arr {
		if v == item {
			return true
		}
	}
	return false
}

func findAndDelete(arr []string, item string) []string {
	for _, v := range arr {
		if v == item {
			slc := make([]string, 0)
			for _, val := range arr {
				if val != item {
					slc = append(slc, val)
				}
			}
			return slc
		}
	}
	return arr
}

func findAndDelete2(arr []string, item string) []string {
	i := 0
	for _, v := range arr {
		if v != item {
			arr[i] = v
			i++
		}
	}
	return arr[:i]
}
