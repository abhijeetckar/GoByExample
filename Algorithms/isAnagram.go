package main

import "fmt"

func main() {
	s := "aabbbb"
	t := "aaaabb"

	res := isAnagram(s, t)
	fmt.Println(res)
	return
}

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	arr := make([]int, 26)

	for i := 0; i < len(t); i++ {
		arr[s[i]-'a']++
		arr[t[i]-'a']--
	}

	for _, v := range arr {
		if v != 0 {
			return false
		}
	}
	return true
}
