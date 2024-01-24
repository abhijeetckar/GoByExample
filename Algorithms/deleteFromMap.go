package main

import "fmt"

func main() {
	sample := make(map[string]int)
	sample["a"] = 1
	sample["b"] = 2
	sample["c"] = 3

	fmt.Println(sample)

	if _, ok := sample[`c`]; ok {
		delete(sample, `c`)
	}

	fmt.Println(sample)

	delete(sample, `c`)
	fmt.Println(sample)

}
