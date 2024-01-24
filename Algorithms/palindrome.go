package main

import (
	"strconv"
)

func main() {
	x := 121

	s := strconv.Itoa(x)

	n := ""
	for i := len(s) - 1; i >= 0; i-- {
		n = n + string(s[i])
	}

	if n == s {
		return true
	}
	return false
}
