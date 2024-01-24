package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	var arr = [5]string{`a`, `b`, `c`, `d`, `e`}

	fmt.Println(arr)

	res, err := json.Marshal(arr)
	if err != nil {
		fmt.Printf("%+v\n", err)
	}
	fmt.Println(string(res))
}
