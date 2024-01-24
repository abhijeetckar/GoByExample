package main

import "fmt"

func main() {
	arr2 := []int{4, 5, 6, 5, 8}
	res := containsDuplicate(arr2)
	fmt.Println(res)
	return
}

func containsDuplicate(nums []int) bool {
	values := make(map[int]bool)
	fmt.Println(values)
	for i := 0; i < len(nums); i++ {
		if values[nums[i]] {
			// Duplicate detected
			return true
		} else {
			values[nums[i]] = true
		}
	}
	return false
}
